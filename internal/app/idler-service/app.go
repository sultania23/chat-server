package idler_service

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	. "github.com/google/uuid"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
	"github.com/sultania23/chat-server/internal/config"
	"github.com/sultania23/chat-server/internal/model/dto"
	mongorepository "github.com/sultania23/chat-server/internal/repository/mongo-repository"
	postgresrepository "github.com/sultania23/chat-server/internal/repository/postgres-repositrory"
	"github.com/sultania23/chat-server/internal/server"
	"github.com/sultania23/chat-server/internal/service"
	"github.com/sultania23/chat-server/internal/transport/gRPC/client"
	"github.com/sultania23/chat-server/internal/transport/http"
	"github.com/sultania23/chat-server/internal/transport/ws"
	"github.com/sultania23/chat-server/pkg/auth"
	"github.com/sultania23/chat-server/pkg/cache"
	"github.com/sultania23/chat-server/pkg/db/mongo"
	"github.com/sultania23/chat-server/pkg/db/postgres"
	"github.com/sultania23/chat-server/pkg/hash"
	"google.golang.org/grpc"
)

// @title        Idler Application
// @version      1.0
// @description  API Server for keep in touch

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization

// Run initializes whole application

func Run(configPath string) {
	fmt.Println(`
 ================================================
 \\\   ######~~#####~~~##~~~~~~#####~~~#####   \\\
  \\\  ~~##~~~~##~~##~~##~~~~~~##~~~~~~##~~##   \\\
   ))) ~~##~~~~##~~##~~##~~~~~~####~~~~#####     )))
  ///  ~~##~~~~##~~##~~##~~~~~~##~~~~~~##~~##   ///
 ///   ######~~#####~~~######~~#####~~~##~~##  ///
 ================================================
	`)

	cfg, err := config.Init(configPath)
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	postgresDB, err := postgres.NewPostgresPool(postgres.Config{
		Host:            cfg.Postgres.Host,
		Port:            cfg.Postgres.Port,
		DB:              cfg.Postgres.DB,
		User:            cfg.Postgres.User,
		Password:        cfg.Postgres.Password,
		MaxConns:        cfg.Postgres.MaxConns,
		MinConns:        cfg.Postgres.MinConns,
		MaxConnLifetime: cfg.Postgres.MaxConnLifetime,
		MaxConnIdleTime: cfg.Postgres.MaxConnIdleTime,
	})
	if err != nil {
		logrus.Fatalf("error initializing postgres: %s", err.Error())
	}

	mongoClient, err := mongo.NewMongoDb(cfg.Mongo)
	if err != nil {
		logrus.Fatalf("error initializing postgres: %s", err.Error())
	}
	mongoDB := mongoClient.Database(cfg.Mongo.DB)

	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)
	tokenManager := auth.NewJWTTokenManager(cfg.Auth.JWT.SigningKey)

	userCache := cache.NewGCache[string, dto.UserDTO](cfg.Cache.Size, cfg.Cache.Expires)

	postgresRepositories := postgresrepository.NewRepositories(postgresDB)
	mongoRepositories := mongorepository.NewRepositories(mongoDB)

	grpcTarget := fmt.Sprintf("%s:%s", cfg.Mail.Host, cfg.Mail.Port)
	grpcConn, err := grpc.Dial(grpcTarget, grpc.WithInsecure())
	grpcClient := client.NewGrpcClient(grpcConn)

	services := service.NewServices(service.ServicesDepends{
		PostgresRepositories: postgresRepositories,
		MongoRepositories:    mongoRepositories,
		Hasher:               hasher,
		TokenManager:         tokenManager,
		TokenTTL:             cfg.Auth.JWT.TokenTTL,
		UserCache:            userCache,
		GrpcClient:           grpcClient,
	})
	httpHandlers := http.NewHandler(services.UserService, tokenManager, services.ConversationService)
	httpServer := server.NewHTTPServer(cfg, httpHandlers.Init(cfg.HTTP))

	go func() {
		if err := httpServer.Run(); err != nil {
			logrus.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	poolCache := cache.NewGCache[UUID, ws.Pool](10, 30*time.Minute)
	wsHandler := ws.NewHandler(cfg.WS, poolCache, services.MessageService, services.ConversationService).InitWSConversations()
	wsServer := server.NewWSServer(cfg, wsHandler)

	go func() {
		if err := wsServer.Run(); err != nil {
			logrus.Errorf("error occurred while running web socket server: %s\n", err.Error())
		}
	}()

	logrus.Print("IDLER facade application has started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("IDLER facade application shutting down")

	if err := httpServer.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on http server shutting down: %s", err.Error())
	}

	if err := wsServer.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on ws server shutting down: %s", err.Error())
	}

	postgresDB.Close()

	if err := mongoClient.Disconnect(context.Background()); err != nil {
		logrus.Errorf("error occured on mongo connection close: %s", err.Error())
	}
}
