package service

import (
    "context"
    "errors"
    "github.com/FordPipatkittikul/backend-challenge/internal/model"
    "github.com/FordPipatkittikul/backend-challenge/internal/repository"
    "github.com/FordPipatkittikul/backend-challenge/internal/utils"
)

type UserService struct {
    Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{Repo: repo}
}

func (s *UserService) Register(ctx context.Context, user *model.User) error {
    hashed, err := utils.HashPassword(user.Password)
    if err != nil {
        return err
    }
    user.Password = hashed
    return s.Repo.CreateUser(ctx, user)
}

func (s *UserService) Login(ctx context.Context, email, password string) (string, error) {
    user, err := s.Repo.GetUserByEmail(ctx, email)
    if err != nil || !utils.CheckPasswordHash(password, user.Password) {
        return "", errors.New("invalid credentials")
    }
    return utils.GenerateJWT(user.Email)
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
    return s.Repo.GetUserByID(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context) ([]model.User, error) {
    return s.Repo.ListUsers(ctx)
}

func (s *UserService) UpdateUser(ctx context.Context, id, name, email string) error {
    return s.Repo.UpdateUser(ctx, id, name, email)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
    return s.Repo.DeleteUser(ctx, id)
}