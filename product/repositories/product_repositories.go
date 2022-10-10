package repositories

import (
	"context"
	"gocommerce/product/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type IProductRepository interface {
	AddProductCtx(ctx context.Context, product *models.Product) error
}

type ProductRepository struct {
	db *mongo.Client
}

func NewProductRepository(db *mongo.Client) IProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) AddProductCtx(ctx context.Context, product *models.Product) error {
	_, err := p.db.Database("gocommerce").Collection("product").InsertOne(ctx, product)
	return err
}
