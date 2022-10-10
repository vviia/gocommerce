package cmd

import (
	"gocommerce/product"

	"go.mongodb.org/mongo-driver/mongo"
)

type Dependencies struct {
	Mongo *mongo.Client

	product product.IProductController
}
