package store

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	goauth "github.com/larrypalm/go-auth"
)

type Store struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context) (*Store, error) {
	pool, err := pgxpool.New(ctx, "postgresql://larry@localhost:5432/larry-golang-testing")
	if err != nil {
		return nil, fmt.Errorf("what")
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return &Store{Pool: pool}, nil
}

func (s *Store) CreateUser(ctx context.Context, email, passwordHash, name string) (goauth.User, error) {
	fmt.Println("CreateUser")
	return goauth.User{}, nil
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (goauth.User, error) {
	fmt.Println("GetUserByEmail")
	return goauth.User{}, nil
}

func (s *Store) GetUserByID(ctx context.Context, id uuid.UUID) (goauth.User, error) {
	fmt.Println("GetUserByID")
	return goauth.User{}, nil
}

func (s *Store) GetUsersByIDs(ctx context.Context, ids []uuid.UUID) ([]goauth.User, error) {
	fmt.Println("GetUsersByIDs")
	return []goauth.User{}, nil
}

func (s *Store) UpdatePassword(ctx context.Context, userID uuid.UUID, passwordHash string) error {
	fmt.Println("UpdatePassword")
	return nil
}

func (s *Store) VerifyEmail(ctx context.Context, userID uuid.UUID) error {
	fmt.Println("VerifyEmail")
	return nil
}

func (s *Store) SaveRefreshToken(ctx context.Context, token goauth.RefreshToken) error {
	fmt.Println("SaveRefreshToken")
	return nil
}
func (s *Store) GetRefreshToken(ctx context.Context, tokenHash string) (goauth.RefreshToken, error) {
	fmt.Println("GetRefreshToken")
	return goauth.RefreshToken{}, nil
}

func (s *Store) RevokeRefreshToken(ctx context.Context, tokenHash string) error {
	fmt.Println("RevokeRefreshToken")
	return nil
}

func (s *Store) RevokeAllRefreshTokens(ctx context.Context, userID uuid.UUID) error {
	fmt.Println("RevokeAllRefreshTokens")
	return nil
}
