package repositories

import (
	"context"
	"fmt"
	"testing"

	"github.com/Pharseus/crud_golang.git/api/config"
	"github.com/Pharseus/crud_golang.git/api/entities"
	repositoryImpl "github.com/Pharseus/crud_golang.git/api/repositories/RepositoryImpl"
)

func TestCreateProduct(t *testing.T) {
	cfg := config.LoadConfig()
	productRepository := repositoryImpl.NewProductRepositoryImpl(config.GetConnection(cfg))
	ctx := context.Background()

	product := entities.Product{
		Name:  "Sedan",
		SKU:   "23k10001",
		Price: 50000,
		Stock: 3,
	}

	result, err := productRepository.CreateProduct(ctx, product)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Id)
}

func TestFindProduct(t *testing.T) {
	cfg := config.LoadConfig()
	repositoryProduct := repositoryImpl.NewProductRepositoryImpl(config.GetConnection(cfg))
	ctx := context.Background()
	var product []entities.Product
	product, id, err := repositoryProduct.FindProduct(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(product, id)

}

func TestFindProductById(t *testing.T) {
	cfg := config.LoadConfig()
	repositoryProduct := repositoryImpl.NewProductRepositoryImpl(config.GetConnection(cfg))
	ctx := context.Background()

	var product entities.Product
	id := 2
	product, err := repositoryProduct.FindProductById(ctx, int32(id))
	if err != nil {
		panic(err)
	}
	fmt.Printf("ambil product dengan id - %d Berhasil", id)
	fmt.Println(product)
}

func TestUpdateProductbyId(t *testing.T) {
	cfg := config.LoadConfig()
	repositoryProduct := repositoryImpl.NewProductRepositoryImpl(config.GetConnection(cfg))
	ctx := context.Background()

	id := 3
	product := entities.Product{
		Name:  "Harley Davidson",
		Price: 650000,
		Stock: 4,
	}
	product, err := repositoryProduct.UpdateProductById(ctx, product, int32(id))
	if err != nil {
		panic(err)
	}
	fmt.Println(product, id)
}

func TestDeleteProduct(t *testing.T) {
	cfg := config.LoadConfig()
	repositoryProduct := repositoryImpl.NewProductRepositoryImpl(config.GetConnection(cfg))
	ctx := context.Background()

	id := 3
	err := repositoryProduct.DeleteProductById(ctx, int32(id))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product dengan id %d Berhasil dihapus", id)
}
