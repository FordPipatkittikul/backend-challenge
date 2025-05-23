# Golang User API

This is a RESTful API built in Go that manages users, using MongoDB for storage and JWT for authentication.

## Features
- User registration, login, CRUD
- JWT authentication (HS256)
- MongoDB integration
- Logging middleware
- Background goroutine to count users
- Clean architecture

## Setup

### Prerequisites
- Go 1.20+
- MongoDB (local or Docker)

### Run MongoDB (optional)
```sh
docker run -d -p 27017:27017 --name mongo mongo
```

### Clone and Run
```sh
git clone https://github.com/your-username/user-api.git
cd user-api
go mod tidy
go run ./cmd
```

### Environment Variables
| Variable         | Default Value        |
|------------------|----------------------|
| `MONGO_URI`      | mongodb://localhost:27017 |
| `MONGO_DB`       | userdb               |
| `MONGO_COLLECTION`| users               |
| `JWT_SECRET`     | supersecretkey       |

### Sample API

**Register:**
```
POST /register
{
  "name": "John",
  "email": "john@example.com",
  "password": "pass123"
}
```

**Login:**
```
POST /login
{
  "email": "john@example.com",
  "password": "pass123"
}
Response: { "token": "<JWT>" }
```

**Protected Route:**
```
GET /users
Headers: Authorization: Bearer <JWT>
```