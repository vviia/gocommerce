package services

import (
	"context"
	"gocommerce/product/dto"
	"gocommerce/product/models"
	"gocommerce/product/repositories"
	"gocommerce/shared/entities"
)

// this layer only for logic purpose

type IProductService interface {
	ProductViewerAdminCtx(ctx context.Context, in *dto.AddProductRequest) (out entities.BaseResposne[string])
}

type ProductService struct {
	ProductRepo repositories.IProductRepository
}

func NewProductService(
	productRepo repositories.IProductRepository,
) IProductService {
	return &ProductService{
		ProductRepo: productRepo,
	}
}

func (p *ProductService) ProductViewerAdminCtx(ctx context.Context, in *dto.AddProductRequest) (out entities.BaseResposne[string]) {

	product := models.Product{
		ID:          generateMongoID(),
		ProductName: in.ProductName,
		Price:       in.Price,
		Rating:      in.Rating,
		Image:       in.Image,
	}
	err := p.ProductRepo.AddProductCtx(ctx, &product)
	if err != nil {
		out.SetCode(500, err)
		return
	}

	out.Message = "Successfully added our Product Admin!!"
	return
}
