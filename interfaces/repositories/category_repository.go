package repositories

import (
	"context"

	"errors"

	"github.com/anderson2093/apiGolang/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoCategoryRepository struct {
	collection *mongo.Collection
}

func NewMongoCategoryRepository(db *mongo.Database) *MongoCategoryRepository {
	return &MongoCategoryRepository{
		collection: db.Collection("categories"),
	}
}

func (r *MongoCategoryRepository) Create(ctx context.Context, category *models.Category) error {
	_, err := r.collection.InsertOne(ctx, category)
	return err
}

func (r *MongoCategoryRepository) GetAll(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var category models.Category
		if err := cursor.Decode(&category); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
func (r *MongoCategoryRepository) GetByID(ctx context.Context, id string) (*models.Category, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID")
	}

	var category models.Category
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&category)
	return &category, err
}

func (r *MongoCategoryRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID")
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
