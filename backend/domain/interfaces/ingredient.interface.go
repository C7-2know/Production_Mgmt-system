package interfaces

import (
	"context"

	"soap_factory/domain/entities"
)

type IngredientService interface {
	CreateIngredient(ctx context.Context, ingredient *entities.Ingredient) error
	UpdateIngredient(ctx context.Context, id string, ingredient *entities.Ingredient) error
	GetIngredient(ctx context.Context, id string) (*entities.Ingredient, error)
	ListIngredients(ctx context.Context) ([]*entities.Ingredient, error)
	RecordPurchase(ctx context.Context, purchase *entities.Purchase) error
	GetStockLevels(ctx context.Context) ([]IngredientStock, error)
	DeleteIngredient(ctx context.Context, id string) error
}

type IngredientStock struct {
	Ingredient *entities.Ingredient
	CurrentStock float64
	MinStock     float64
	NeedsReorder bool
}

type IngredientRepositoryInterface interface {
    Create(ctx context.Context, ingredient *entities.Ingredient) error
    Update(ctx context.Context, id string, ingredient *entities.Ingredient) error
    FindByID(ctx context.Context, id string) (*entities.Ingredient, error)
    FindByName(ctx context.Context, name string) (*entities.Ingredient, error)
    FindAll(ctx context.Context) ([]*entities.Ingredient, error)
    UpdateStock(ctx context.Context, ingredientID string, quantity float64) error
	Delete(ctx context.Context, id string) error
}