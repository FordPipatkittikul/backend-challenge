package repository

import (
    "context"
    "errors"
    "time"

    "github.com/FordPipatkittikul/backend-challenge/internal/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type mongoUserRepository struct {
    db *MongoDB
}

func NewUserRepository(db *MongoDB) UserRepository {
    return &mongoUserRepository{db: db}
}

func (r *mongoUserRepository) CreateUser(ctx context.Context, user *model.User) error {
    user.CreatedAt = time.Now()
    _, err := r.db.Collection.InsertOne(ctx, user)
    return err
}

func (r *mongoUserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
    var user model.User
    err := r.db.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
    return &user, err
}

func (r *mongoUserRepository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    var user model.User
    err = r.db.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
    return &user, err
}

func (r *mongoUserRepository) ListUsers(ctx context.Context) ([]model.User, error) {
    cur, err := r.db.Collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cur.Close(ctx)

    var users []model.User
    for cur.Next(ctx) {
        var user model.User
        err := cur.Decode(&user)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}

func (r *mongoUserRepository) UpdateUser(ctx context.Context, id string, name string, email string) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    update := bson.M{
        "$set": bson.M{
            "name":  name,
            "email": email,
        },
    }

    res, err := r.db.Collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
    if res.MatchedCount == 0 {
        return errors.New("user not found")
    }
    return err
}

func (r *mongoUserRepository) DeleteUser(ctx context.Context, id string) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    _, err = r.db.Collection.DeleteOne(ctx, bson.M{"_id": objID})
    return err
}
