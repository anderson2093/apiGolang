package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id" example:"67791e1d1b8af760b0149bb5"`
	Name        string             `bson:"name" json:"name" binding:"required" example:"Pantalones"`
	Description string             `bson:"description" json:"description" example:"Categoría de pantalones de vestir y casual"`
}
