package model

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type User struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name      string             `bson:"name" json:"name"`
    Email     string             `bson:"email" json:"email"`
    Password  string             `bson:"password" json:"-"`
    CreatedAt time.Time          `bson:"createdAt" json:"created_at"`
}

