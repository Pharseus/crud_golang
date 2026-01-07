package route

import (
	"net/http"

	"github.com/Pharseus/crud_golang.git/api/controllers"
	"github.com/Pharseus/crud_golang.git/api/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func StartRouting(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API berhasil di jalankan"))
	})

	// api buat user
	r.Route("/v1", func(r chi.Router) {
		userService := services.NewUserService(db)
		userController := controllers.NewUserController(userService)

		r.Route("/users", func(r chi.Router) {
			r.Post("/", userController.Create)
			r.Get("/", userController.GetAll)
			r.Get("/{id}", userController.GetById)
			r.Put("/{id}", userController.Update)
			r.Delete("/{id}", userController.Delete)
		})
	})

	return r
}
