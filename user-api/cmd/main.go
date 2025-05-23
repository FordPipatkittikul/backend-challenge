package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"

    "github.com/gorilla/mux"
    "github.com/FordPipatkittikul/backend-challenge/config"
    "github.com/FordPipatkittikul/backend-challenge/internal/middleware"
    "github.com/FordPipatkittikul/backend-challenge/internal/model"
    "github.com/FordPipatkittikul/backend-challenge/internal/repository"
    "github.com/FordPipatkittikul/backend-challenge/internal/service"
)

func main() {
    // Setup DB
    db, err := repository.NewMongoDB(config.MongoURI, config.MongoDBName, config.MongoCollName)
    if err != nil {
        log.Fatal(err)
    }
    repo := repository.NewUserRepository(db)
    svc := service.NewUserService(repo)

    // Router and Middleware
    r := mux.NewRouter()
    r.Use(middleware.LoggingMiddleware)

    // Routes
    r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
        var user model.User
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        if err := svc.Register(r.Context(), &user); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
    }).Methods("POST")

    r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        var creds struct {
            Email    string `json:"email"`
            Password string `json:"password"`
        }
        if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        token, err := svc.Login(r.Context(), creds.Email, creds.Password)
        if err != nil {
            http.Error(w, err.Error(), http.StatusUnauthorized)
            return
        }
        json.NewEncoder(w).Encode(map[string]string{"token": token})
    }).Methods("POST")

    r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        users, err := svc.ListUsers(r.Context())
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(users)
    }).Methods("GET")

    // Background goroutine
    go func() {
        for {
            time.Sleep(10 * time.Second)
            users, err := svc.ListUsers(context.Background())
            if err == nil {
                log.Printf("[INFO] User count: %d", len(users))
            }
        }
    }()

    // Graceful shutdown
    srv := &http.Server{
        Addr:    ":8080",
        Handler: r,
    }

    go func() {
        log.Println("Server running on :8080")
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
    }()

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c
    log.Println("Shutting down...")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    srv.Shutdown(ctx)
    log.Println("Server gracefully stopped")
}