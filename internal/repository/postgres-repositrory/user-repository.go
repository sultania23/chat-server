package postgres_repository

import (
	"context"
	"fmt"
	. "github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/tuxoo/idler/internal/model/dto"
	"github.com/tuxoo/idler/internal/model/entity"
)

type UserRepository struct {
	//db *sqlx.DB
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(ctx context.Context, user entity.User) (*dto.UserDTO, error) {
	var newUser dto.UserDTO
	query := fmt.Sprintf("INSERT INTO %s (name, login_email, password_hash, registered_at, visited_at, role) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, login_email, registered_at, visited_at, role, is_enabled", userTable)
	row := r.db.QueryRow(ctx, query, user.Name, user.LoginEmail, user.PasswordHash, user.RegisteredAt, user.VisitedAt, user.Role)

	if err := row.Scan(&newUser.Id, &newUser.Name, &newUser.LoginEmail, &newUser.RegisteredAt, &newUser.VisitedAt, &newUser.Role, &newUser.IsEnabled); err != nil {
		return &newUser, err
	}

	return &newUser, nil
}

func (r *UserRepository) UpdateById(ctx context.Context, id UUID) error {
	query := fmt.Sprintf("UPDATE %s SET is_enabled=true WHERE id=$1", userTable)
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *UserRepository) FindByCredentials(ctx context.Context, email, password string) (*dto.UserDTO, error) {
	var user dto.UserDTO
	query := fmt.Sprintf("SELECT id, name, login_email, registered_at, visited_at, role, is_enabled FROM %s WHERE is_enabled=true AND login_email=$1 AND password_hash=$2", userTable)
	row := r.db.QueryRow(ctx, query, email, password)

	if err := row.Scan(&user.Id, &user.Name, &user.LoginEmail, &user.RegisteredAt, &user.VisitedAt, &user.Role, &user.IsEnabled); err != nil {
		return &user, err
	}

	return &user, nil
}

func (r *UserRepository) FindById(ctx context.Context, id UUID) (*dto.UserDTO, error) {
	var user dto.UserDTO
	query := fmt.Sprintf("SELECT id, name, login_email, registered_at, visited_at, role, is_enabled FROM %s WHERE is_enabled=true AND id=$1", userTable)
	row := r.db.QueryRow(ctx, query, id)

	if err := row.Scan(&user.Id, &user.Name, &user.LoginEmail, &user.RegisteredAt, &user.VisitedAt, &user.Role, &user.IsEnabled); err != nil {
		return &user, err
	}

	return &user, nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]dto.UserDTO, error) {
	var users []dto.UserDTO
	//query := fmt.Sprintf("SELECT id, name, login_email, registered_at, visited_at, role, is_enabled FROM %s WHERE is_enabled=true", userTable)
	//if err := r.db.Select(&users, query); err != nil {
	//	return nil, err
	//}

	return users, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string, isEnabled bool) (*dto.UserDTO, error) {
	var user dto.UserDTO
	query := fmt.Sprintf("SELECT id, name, login_email, registered_at, visited_at, role, is_enabled FROM %s WHERE login_email=$1 AND is_enabled=$2", userTable)
	row := r.db.QueryRow(ctx, query, email, isEnabled)

	if err := row.Scan(&user.Id, &user.Name, &user.LoginEmail, &user.RegisteredAt, &user.VisitedAt, &user.Role, &user.IsEnabled); err != nil {
		return &user, err
	}

	return &user, nil
}
