package repositories

import (
    "context"
	"errors"
	"time"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"

    "soap_factory/domain/entities"
)

type IngredientRepository struct {
    collection *mongo.Collection
}

func NewIngredientRepository(db *mongo.Database) *IngredientRepository {
    return &IngredientRepository{
        collection: db.Collection("ingredients"),
    }
}

func (r *IngredientRepository) Create(ctx context.Context, ingredient *entities.Ingredient) error {
	if ingredient.ID == "" {
		ingredient.ID = primitive.NewObjectID().Hex()
	}

	_, err := r.collection.InsertOne(ctx, ingredient)
	return err
}

func (r *IngredientRepository) FindAll(ctx context.Context) ([]*entities.Ingredient, error) {
	var ingredients []*entities.Ingredient
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var ingredient entities.Ingredient
		if err := cursor.Decode(&ingredient); err != nil {
			return nil, err
		}
		ingredients = append(ingredients, &ingredient)
	}

	return ingredients, nil
}

func (r *IngredientRepository) FindByID(ctx context.Context, id string) (*entities.Ingredient, error) {
	var ingredient entities.Ingredient
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&ingredient)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &ingredient, nil
}

func (r *IngredientRepository) Update(ctx context.Context, id string, updatedData *entities.Ingredient) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": updatedData},
	)
	return err
}

func (r *IngredientRepository) Delete(ctx context.Context, id string) error {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no document found to delete")
	}
	return nil
}

func (r *IngredientRepository) FindByName(ctx context.Context, name string) (*entities.Ingredient, error) {
	var ingredient entities.Ingredient
	err := r.collection.FindOne(ctx, bson.M{"name": name}).Decode(&ingredient)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &ingredient, nil
}

func (r *IngredientRepository) UpdateStock(ctx context.Context, ingredientID string, quantity float64) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": ingredientID},
		bson.M{"$inc": bson.M{"current_stock": quantity}, "$set": bson.M{"updated_at": time.Now()}},
	)
	return err
}