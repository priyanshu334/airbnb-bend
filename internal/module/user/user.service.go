package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var ErrEmailAlreadyExists = errors.New("email already exists")
var ErrInvalidCredentials = errors.New("invalid credentials")

type Service interface {
	Register(ctx context.Context, input RegisterInput) (*User, error)
	GetById(ctx context.Context, id uuid.UUID) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

type RegisterInput struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
	Role      Role
}

func (s *service) Register(ctx context.Context, input RegisterInput) (*User, error) {
	_, err := s.repo.FindByEmail(ctx, input.Email)

	if err != nil {
		return nil, ErrEmailAlreadyExists
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return nil, err
	}
	user := &User{
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Password:  string(hash),
		Role:      input.Role,
	}
	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, err
}

func (s *service) GetById(ctx context.Context, id uuid.UUID) (*User, error) {
	return s.repo.FindById(ctx, id)
}
