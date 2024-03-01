package service

import (
	"context"
	"errors"
	. "github.com/google/uuid"
	"github.com/tuxoo/idler/internal/model/dto"
	"github.com/tuxoo/idler/internal/model/entity"
	postgres_repository "github.com/tuxoo/idler/internal/repository/postgres-repositrory"
	"github.com/tuxoo/idler/internal/transport/gRPC/client"
	"github.com/tuxoo/idler/pkg/auth"
	"github.com/tuxoo/idler/pkg/cache"
	"github.com/tuxoo/idler/pkg/hash"
	"time"
)

type UserService struct {
	repository   postgres_repository.Users
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager
	tokenTTL     time.Duration
	userCache    cache.Cache[string, dto.UserDTO]
	grpcClient   *client.GrpcClient
}

func NewUserService(repository postgres_repository.Users, hasher hash.PasswordHasher, tokenManager auth.TokenManager, tokenTTL time.Duration, userCache cache.Cache[string, dto.UserDTO], grpcClient *client.GrpcClient) *UserService {
	return &UserService{
		repository:   repository,
		hasher:       hasher,
		tokenManager: tokenManager,
		tokenTTL:     tokenTTL,
		userCache:    userCache,
		grpcClient:   grpcClient,
	}
}

func (s *UserService) SignUp(ctx context.Context, dto dto.SignUpDTO) error {
	user := entity.User{
		Name:         dto.Name,
		LoginEmail:   dto.Email,
		PasswordHash: s.hasher.Hash(dto.Password),
		RegisteredAt: time.Now(),
		VisitedAt:    time.Now(),
		Role:         entity.UserRole,
	}

	_, err := s.repository.Save(ctx, user)
	if err != nil {
		return err
	}

	//_, err = s.grpcClient.MailSender.SendMail(ctx, &api.Mail{
	//	Address: dto.Email,
	//})
	//if err != nil {
	//	return err
	//}

	return nil
}

func (s *UserService) VerifyUser(ctx context.Context, verifyDTO dto.VerifyDTO) error {
	user, err := s.repository.FindByEmail(ctx, verifyDTO.Email, false)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("unknown user")
	}

	if user.IsEnabled {
		return errors.New("user already active")
	}

	if verifyDTO.CheckCode != s.hasher.Hash(user.Name) {
		return errors.New("illegal check code")
	}

	return s.repository.UpdateById(ctx, user.Id)
}

func (s *UserService) SignIn(ctx context.Context, dto dto.SignInDTO) (token auth.Token, err error) {
	user, err := s.repository.FindByCredentials(ctx, dto.Email, s.hasher.Hash(dto.Password))
	if err != nil {
		return "", err
	}

	id := user.Id.String()
	if err := s.userCache.Set(ctx, id, user); err != nil {
		return token, err
	}

	token, err = s.tokenManager.GenerateToken(id, s.tokenTTL)
	return
}

func (s *UserService) GetById(ctx context.Context, id UUID) (user *dto.UserDTO, err error) {
	user, err = s.userCache.Get(ctx, id.String())
	if user == nil && err != nil {
		user, err = s.repository.FindById(ctx, id)
	}
	return
}

func (s *UserService) GetAll(ctx context.Context) ([]dto.UserDTO, error) {
	return s.repository.FindAll(ctx)
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*dto.UserDTO, error) {
	return s.repository.FindByEmail(ctx, email, true)
}
