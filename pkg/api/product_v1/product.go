package product_v1

import (
	"context"
	"github.com/iamgadfly/go-echo-api/config"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/iamgadfly/go-echo-api/internal/products"
	"github.com/iamgadfly/go-echo-api/internal/products/repository"
	product "github.com/iamgadfly/go-echo-api/pkg/api/product_v1/proto"
)

// ProductServer ..
type ProductServer struct {
	product.UnimplementedProductServer
	cfg     *config.Config
	useCase products.ProductUseCase
	repo    *repository.ProductRepo
}

// NewProductServer ..
func NewProductServer(cfg *config.Config, repo *repository.ProductRepo, uc products.ProductUseCase) *ProductServer {
	return &ProductServer{
		useCase: uc,
		cfg:     cfg,
		repo:    repo,
	}
}

// Create ..
func (p ProductServer) Create(ctx context.Context, req *product.CreateRequest) (*product.ProductResponse, error) {
	prod, err := p.useCase.ParseByLink(req.GetLink())
	if err != nil {
		return nil, err
	}

	return &product.ProductResponse{Data: p.createResponse(&prod)}, nil
}

// GetById ..
func (p ProductServer) GetById(ctx context.Context, req *product.GetByIdRequest) (*product.ProductResponse, error) {
	prod, err := p.repo.GetById(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}

	return &product.ProductResponse{Data: p.createResponse(prod)}, nil
}

// Search ..
func (p ProductServer) Search(ctx context.Context, req *product.SearchRequest) (*product.ProductListResponse, error) {
	products, err := p.repo.Search(ctx, req.GetWord())
	if err != nil {
		return nil, err
	}

	prods := make([]*product.ProductData, len(products))
	for i, v := range products {
		prods[i] = p.createResponse(&v)
	}

	return &product.ProductListResponse{Data: prods}, nil
}

// createResponse ..
func (p ProductServer) createResponse(prod *models.Product) *product.ProductData {
	return &product.ProductData{
		Id:          uint64(prod.ID),
		Name:        prod.Name,
		Price:       uint64(prod.Price),
		SalePrice:   uint64(prod.SalePrice),
		Description: prod.Description,
		Image:       prod.Image,
		ShopId:      int64(prod.ShopId),
		Address:     prod.Address,
		Color:       prod.Color,
		Link:        prod.Link,
	}
}
