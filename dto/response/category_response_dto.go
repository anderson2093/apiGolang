package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type CategoryResponseDto struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
}
