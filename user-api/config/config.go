package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, relying on environment variables")
    }
}
var (
    MongoURI      = getEnv("MONGO_URI", "mongodb://localhost:27017")
    MongoDBName   = getEnv("MONGO_DB", "userdb")
    MongoCollName = getEnv("MONGO_COLLECTION", "users")
    JWTSecret     = getEnv("JWT_SECRET", "supersecretkey")
)

func getEnv(key, defaultVal string) string {
    if val, ok := os.LookupEnv(key); ok {
        return val
    }
    return defaultVal
}