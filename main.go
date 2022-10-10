package main

import (
	"context"
	"log"

	"gocommerce/product"
	productRepository "gocommerce/product/repositories"
	productService "gocommerce/product/services"
	"gocommerce/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()

	// intitialize database
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// initialize repositories
	productRepositoryImpl := productRepository.NewProductRepository(mongoClient)

	// intialize services
	productServiceImpl := productService.NewProductService(productRepositoryImpl)

	// initialze controller
	productControllerImpl := product.NewProductController(productServiceImpl)

	routes.ProductRoutes(r, productControllerImpl)

	r.Run(":8080")
}
