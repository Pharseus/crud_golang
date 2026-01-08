package route

import (
	"net/http"

	"github.com/Pharseus/crud_golang.git/api/controllers"
	"github.com/Pharseus/crud_golang.git/api/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"

	_ "github.com/Pharseus/crud_golang.git/docs"
	httpSwagger "github.com/swaggo/http-swagger"
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

	r.Route("/v1", func(r chi.Router) {
		userService := services.NewUserService(db)
		userController := controllers.NewUserController(userService)

		// api buat user
		r.Route("/users", func(r chi.Router) {
			r.Post("/", userController.Create)
			r.Get("/", userController.GetAll)
			r.Get("/{id}", userController.GetById)
			r.Put("/{id}", userController.Update)
			r.Delete("/{id}", userController.Delete)
		})

		// product
		productService := services.NewProductService(db)
		productController := controllers.NewProductController(productService)

		r.Route("/products", func(r chi.Router) {
			r.Post("/", productController.Create)       // POST /v1/products
			r.Get("/", productController.GetAll)        // GET /v1/products
			r.Get("/{id}", productController.GetById)   // GET /v1/products/{id}
			r.Put("/{id}", productController.Update)    // PUT /v1/products/{id}
			r.Delete("/{id}", productController.Delete) // DELETE /v1/products/{id}
		})

	})

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"),
	))

	return r
}
