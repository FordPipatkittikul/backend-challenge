package repository

import (
    "context"

    "github.com/FordPipatkittikul/backend-challenge/internal/model"
)

type UserRepository interface {
    CreateUser(ctx context.Context, user *model.User) error
    GetUserByEmail(ctx context.Context, email string) (*model.User, error)
    GetUserByID(ctx context.Context, id string) (*model.User, error)
    ListUsers(ctx context.Context) ([]model.User, error)
    UpdateUser(ctx context.Context, id string, name string, email string) error
    DeleteUser(ctx context.Context, id string) error
}
