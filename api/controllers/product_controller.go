package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Pharseus/crud_golang.git/api/helper"
	"github.com/Pharseus/crud_golang.git/api/payloads"
	"github.com/Pharseus/crud_golang.git/api/services"
	"github.com/go-chi/chi/v5"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

// Create godoc
// @Summary      Create a new product
// @Description  Create a new product with name, SKU, price, and stock
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request body payloads.CreateProductRequest true "Create Product Request"
// @Success      201 {object} payloads.SuccessResponse{data=payloads.ProductResponse}
// @Failure      400 {object} payloads.ErrorResponse
// @Failure      500 {object} payloads.ErrorResponse
// @Router       /v1/products [post]
func (c *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	var req payloads.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.RespondError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	product, err := c.productService.Create(r.Context(), req)
	if err != nil {
		helper.RespondError(w, http.StatusInternalServerError, "Failed to create product", err)
		return
	}

	helper.RespondSuccess(w, http.StatusCreated, "Product created successfully", product)
}

// GetAll godoc
// @Summary      Get all products
// @Description  Get list of all products
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200 {object} payloads.SuccessResponse{data=[]payloads.ProductResponse}
// @Failure      500 {object} payloads.ErrorResponse
// @Router       /v1/products [get]
func (c *ProductController) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := c.productService.FindAll(r.Context())
	if err != nil {
		helper.RespondError(w, http.StatusInternalServerError, "Failed to get products", err)
		return
	}

	helper.RespondSuccess(w, http.StatusOK, "Products retrieved successfully", products)
}

// GetById godoc
// @Summary      Get product by ID
// @Description  Get a single product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id path int true "Product ID"
// @Success      200 {object} payloads.SuccessResponse{data=payloads.ProductResponse}
// @Failure      400 {object} payloads.ErrorResponse
// @Failure      404 {object} payloads.ErrorResponse
// @Router       /v1/products/{id} [get]
func (c *ProductController) GetById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		helper.RespondError(w, http.StatusBadRequest, "Invalid product ID", err)
		return
	}

	product, err := c.productService.FindById(r.Context(), int32(id))
	if err != nil {
		helper.RespondError(w, http.StatusNotFound, "Product not found", err)
		return
	}

	helper.RespondSuccess(w, http.StatusOK, "Product retrieved successfully", product)
}

// Update godoc
// @Summary      Update product
// @Description  Update product information by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id path int true "Product ID"
// @Param        request body payloads.UpdateProductRequest true "Update Product Request"
// @Success      200 {object} payloads.SuccessResponse{data=payloads.ProductResponse}
// @Failure      400 {object} payloads.ErrorResponse
// @Failure      500 {object} payloads.ErrorResponse
// @Router       /v1/products/{id} [put]
func (c *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		helper.RespondError(w, http.StatusBadRequest, "Invalid product ID", err)
		return
	}

	var req payloads.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.RespondError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	product, err := c.productService.Update(r.Context(), int32(id), req)
	if err != nil {
		helper.RespondError(w, http.StatusInternalServerError, "Failed to update product", err)
		return
	}

	helper.RespondSuccess(w, http.StatusOK, "Product updated successfully", product)
}

// Delete godoc
// @Summary      Delete product
// @Description  Delete product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id path int true "Product ID"
// @Success      200 {object} payloads.SuccessResponse
// @Failure      400 {object} payloads.ErrorResponse
// @Failure      500 {object} payloads.ErrorResponse
// @Router       /v1/products/{id} [delete]
func (c *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		helper.RespondError(w, http.StatusBadRequest, "Invalid product ID", err)
		return
	}

	if err := c.productService.Delete(r.Context(), int32(id)); err != nil {
		helper.RespondError(w, http.StatusInternalServerError, "Failed to delete product", err)
		return
	}

	helper.RespondSuccess(w, http.StatusOK, "Product deleted successfully", nil)
}
