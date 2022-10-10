package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID
	ProductName string
	Price       float64
	Rating      float64
	Image       string
}
