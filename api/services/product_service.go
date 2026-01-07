package services

import (
	"context"

	"github.com/Pharseus/crud_golang.git/api/entities"
	"github.com/Pharseus/crud_golang.git/api/payloads"
	repository "github.com/Pharseus/crud_golang.git/api/repositories/Repository"
	repositoryImpl "github.com/Pharseus/crud_golang.git/api/repositories/RepositoryImpl"
	"gorm.io/gorm"
)

type ProductService interface {
	Create(ctx context.Context, req payloads.CreateProductRequest) (payloads.ProductResponse, error)
	FindAll(ctx context.Context) ([]payloads.ProductResponse, error)
	FindById(ctx context.Context, id int32) (payloads.ProductResponse, error)
	Update(ctx context.Context, id int32, req payloads.UpdateProductRequest) (payloads.ProductResponse, error)
	Delete(ctx context.Context, id int32) error
}

type productServiceImpl struct {
	prodRepo repository.ProductRepository
}

func NewProductService(db *gorm.DB) ProductService {
	return &productServiceImpl{
		prodRepo: repositoryImpl.NewProductRepositoryImpl(db),
	}
}

func (s *productServiceImpl) Create(ctx context.Context, req payloads.CreateProductRequest) (payloads.ProductResponse, error) {
	product := entities.Product{
		Name:  req.Name,
		SKU:   req.SKU,
		Price: req.Price,
		Stock: req.Stock,
	}

	created, err := s.prodRepo.CreateProduct(ctx, product)
	if err != nil {
		return payloads.ProductResponse{}, err
	}

	return payloads.ProductResponse{
		Id:        created.Id,
		Name:      created.Name,
		SKU:       created.SKU,
		Price:     created.Price,
		Stock:     created.Stock,
		CreatedAt: created.CreatedAt,
		UpdatedAt: created.UpdatedAt,
	}, nil
}

func (s *productServiceImpl) FindAll(ctx context.Context) ([]payloads.ProductResponse, error) {
	products, _, err := s.prodRepo.FindProduct(ctx)
	if err != nil {
		return nil, err
	}

	var responses []payloads.ProductResponse
	for _, product := range products {
		responses = append(responses, payloads.ProductResponse{
			Id:        product.Id,
			Name:      product.Name,
			SKU:       product.SKU,
			Price:     product.Price,
			Stock:     product.Stock,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		})
	}

	return responses, nil
}

func (s *productServiceImpl) FindById(ctx context.Context, id int32) (payloads.ProductResponse, error) {
	product, err := s.prodRepo.FindProductById(ctx, id)
	if err != nil {
		return payloads.ProductResponse{}, err
	}

	return payloads.ProductResponse{
		Id:        product.Id,
		Name:      product.Name,
		SKU:       product.SKU,
		Price:     product.Price,
		Stock:     product.Stock,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}, nil
}

func (s *productServiceImpl) Update(ctx context.Context, id int32, req payloads.UpdateProductRequest) (payloads.ProductResponse, error) {
	product := entities.Product{
		Name: req.Name,
		SKU:  req.SKU,
	}

	if req.Price != nil {
		product.Price = *req.Price
	}

	if req.Stock != nil {
		product.Stock = *req.Stock
	}

	updated, err := s.prodRepo.UpdateProductById(ctx, product, id)
	if err != nil {
		return payloads.ProductResponse{}, err
	}

	return payloads.ProductResponse{
		Id:        updated.Id,
		Name:      updated.Name,
		SKU:       updated.SKU,
		Price:     updated.Price,
		Stock:     updated.Stock,
		CreatedAt: updated.CreatedAt,
		UpdatedAt: updated.UpdatedAt,
	}, nil
}

func (s *productServiceImpl) Delete(ctx context.Context, id int32) error {
	return s.prodRepo.DeleteProductById(ctx, id)
}
