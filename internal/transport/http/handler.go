package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sultania23/chat-server/internal/config"
	"github.com/sultania23/chat-server/internal/service"
	"github.com/sultania23/chat-server/pkg/auth"
	"net/http"
	"time"
)

type Handler struct {
	userService         service.Users
	tokenManager        auth.TokenManager
	conversationService service.Conversations
}

func NewHandler(userService service.Users, tokenManager auth.TokenManager, dialogService service.Conversations) *Handler {
	return &Handler{
		userService:         userService,
		tokenManager:        tokenManager,
		conversationService: dialogService,
	}
}

func (h *Handler) Init(cfg config.HTTPConfig) *gin.Engine {
	router := gin.New()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		cors.New(corsConfig),
	)

	//docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("api/v1/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	h.initMetrics(router)
	h.initApi(router)

	return router
}

func (h *Handler) initApi(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		h.initUserRoutes(api)
		h.initConversationRoutes(api)
	}
}

func (h *Handler) initMetrics(router *gin.Engine) {
	metrics := router.Group("/metrics")
	{
		h.initMetricRoutes(metrics)
	}
}
