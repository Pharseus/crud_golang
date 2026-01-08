# CRUD API SEDERHANA

REST API menggunakan Go, Chi Router, GORM, dan MySQL.

##  Persyaratan

- Go 1.21 atau lebih tinggi
- MySQL 8.0 atau lebih tinggi
- Postman (untuk testing API)

##  Dependencies

```go
require (
    github.com/go-chi/chi/v5 v5.0.11
    github.com/spf13/viper v1.18.2
    github.com/swaggo/http-swagger v1.3.4
    github.com/swaggo/swag v1.16.3
    golang.org/x/crypto v0.17.0
    gorm.io/driver/mysql v1.5.2
    gorm.io/gorm v1.25.5
)
```

##  Instalasi

### 1. Clone Repository

```bash
git clone https://github.com/Pharseus/crud_golang.git
cd crud_api
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Konfigurasi Environment

env database
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=crud_db

### 4. Buat Database

```sql
CREATE DATABASE crud_db;
```

### 5. Generate Swagger Documentation

```bash
swag init
```

### 6. Jalankan Aplikasi

```bash
go run main.go
```


## 📚 API Documentation

Akses Swagger UI di browser:

**http://localhost:3000/swagger/index.html**

## 📖 Daftar Endpoint

### User API

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/v1/users` | Buat user baru |
| GET | `/v1/users` | Ambil semua user (dengan pagination) |
| GET | `/v1/users/{id}` | Ambil user berdasarkan ID |
| PUT | `/v1/users/{id}` | Update user berdasarkan ID |
| DELETE | `/v1/users/{id}` | Hapus user (soft delete) |

### Product API

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/v1/products` | Buat product baru |
| GET | `/v1/products` | Ambil semua product (dengan pagination) |
| GET | `/v1/products/{id}` | Ambil product berdasarkan ID |
| PUT | `/v1/products/{id}` | Update product berdasarkan ID |
| DELETE | `/v1/products/{id}` | Hapus product (soft delete) |

##  Struktur Project

```
crud_api/
├── main.go                    # Entry point aplikasi
├── .env                       # Konfigurasi environment
├── README.md
├── docs/                      # Swagger documentation (auto-generated)
├── api/
│   ├── config/               # Database & configuration
│   │   ├── loader.go
│   │   ├── app_config.go
│   │   └── database.go
│   ├── entities/             # Database models (GORM)
│   │   ├── user.go
│   │   ├── product.go
│   │   ├── order.go
│   │   └── payment.go
│   ├── payloads/             # Request/Response DTOs
│   │   ├── user_request.go
│   │   ├── user_response.go
│   │   ├── product_request.go
│   │   ├── product_response.go
│   │   └── common.go
│   ├── repositories/         # Data access layer
│   │   ├── Repository/       # Interfaces
│   │   └── RepositoryImpl/   # Implementations
│   ├── services/             # Business logic layer
│   │   ├── user_service.go
│   │   └── product_service.go
│   ├── controllers/          # HTTP handlers
│   │   ├── user_controller.go
│   │   └── product_controller.go
│   ├── helper/               # Utilities
│   │   └── response.go
│   ├── securities/           # Security utilities
│   │   └── password.go
│   ├── middlewares/          # HTTP middlewares
│   │   └── Auth.go
│   └── route/                # Router configuration
│       └── Route.go
```

