package main

import (
	"fmt"
	"net/http"

	"github.com/Pharseus/crud_golang.git/api/config"
	"github.com/Pharseus/crud_golang.git/api/entities"
	"github.com/Pharseus/crud_golang.git/api/route"
)

// @title           CRUD API Documentation
// @version         1.0
// @description     REST API for User and Product Management
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:3000
// @BasePath  /

// @schemes http https
func main() {
	cfg := config.LoadConfig()
	fmt.Println(cfg)
	db := config.GetConnection(cfg)

	err := db.AutoMigrate(
		&entities.User{},
		&entities.Product{},
		&entities.Order{},
		&entities.Payment{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Migrasi berhasil")
	// route.StartRouting(db)

	router := route.StartRouting(db)

	// Start server
	port := ":3000"
	fmt.Printf("🚀 Server running on http://localhost%s\n", port)
	fmt.Printf("📚 Swagger UI: http://localhost%s/swagger/index.html\n", port)
	if err := http.ListenAndServe(port, router); err != nil {
		panic(err)
	}

}
