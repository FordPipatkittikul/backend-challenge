package utils

import (
    "github.com/golang-jwt/jwt/v5"
    "time"
    "github.com/FordPipatkittikul/backend-challenge/config"
)

type JWTClaims struct {
    Email string `json:"email"`
    jwt.RegisteredClaims
}

func GenerateJWT(email string) (string, error) {
    claims := JWTClaims{
        Email: email,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(config.JWTSecret))
}

func ValidateJWT(tokenString string) (*JWTClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(config.JWTSecret), nil
    })
    if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
        return claims, nil
    } else {
        return nil, err
    }
}