package tests

import (
	"context"
	"github.com/iamgadfly/go-echo-api/internal/models"
	"github.com/iamgadfly/go-echo-api/internal/products/usecase"
	"github.com/iamgadfly/go-echo-api/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UnitTestProduct struct {
	suite.Suite
}

func TestUnitTestProduct(t *testing.T) {
	suite.Run(t, &UnitTestProduct{})
}

func (uts *UnitTestProduct) TestParseByLink() {
	repo := mocks.ProductRepository{}
	repo.On("SearchByShopId", mock.AnythingOfType("models.Product")).Return(nil).Once()
	useCase := usecase.NewProductUseCase(nil, &repo, nil)
	_, err := useCase.ParseByLink("https://www.wildberries.ru/catalog/19252625/detail.aspx")
	//uts.T().Log(prod, err)
	uts.Nil(err)
}

func (uts *UnitTestProduct) TestSearch() {
	repo := mocks.ProductRepository{}
	ctx := context.Background()
	repo.On("Search", ctx, mock.Anything).Return([]models.Product{}, nil).Once()
	UC := usecase.NewProductUseCase(nil, &repo, nil)
	p, err := UC.Search(ctx, "Iphone")
	uts.T().Log(p)
	uts.Nil(err)
}
func (uts *UnitTestProduct) TestGetById() {
	repo := mocks.ProductRepository{}
	ctx := context.Background()
	repo.On("GetById", ctx, mock.Anything).Return(&models.Product{}, nil).Once()
	UC := usecase.NewProductUseCase(nil, &repo, nil)
	p, err := UC.GetById(ctx, "322")
	uts.T().Log(p)
	uts.Nil(err)
}
