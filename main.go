package main

import (
	"fmt"
	"net/http"

	"github.com/Pharseus/crud_golang.git/API/config"
	"github.com/Pharseus/crud_golang.git/api/entities"
	"github.com/Pharseus/crud_golang.git/api/route"
)

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
	fmt.Printf("Server running on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, router); err != nil {
		panic(err)
	}

}
