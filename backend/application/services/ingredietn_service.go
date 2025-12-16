package services

import (
	"context"
	"errors"
	"time"
	
    "soap_factory/domain/entities"
	"soap_factory/domain/interfaces"
)

type IngredientService struct {
    repo interfaces.IngredientRepositoryInterface
}

func NewIngredientService(repo interfaces.IngredientRepositoryInterface) *IngredientService {
    return &IngredientService{repo: repo}
}


func (s *IngredientService) CreateIngredient(ctx context.Context, ingredient *entities.Ingredient) error {
	existing, err := s.repo.FindByName(ctx, ingredient.Name)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("ingredient with this name already exists")
	}

	now := time.Now()
	ingredient.CreatedAt = now
	ingredient.UpdatedAt = now

    return s.repo.Create(ctx, ingredient)
}	

func (s *IngredientService) ListIngredients(ctx context.Context) ([]*entities.Ingredient, error) {
	return s.repo.FindAll(ctx)
}

func (s *IngredientService) GetIngredient(ctx context.Context, id string) (*entities.Ingredient, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *IngredientService) UpdateIngredient(ctx context.Context, id string, updatedData *entities.Ingredient) error {
	updatedData.UpdatedAt = time.Now()
	return s.repo.Update(ctx, id, updatedData)
}

func (s *IngredientService) DeleteIngredient(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

// RecordPurchase records a purchase and updates stock
func (s *IngredientService) RecordPurchase(ctx context.Context, purchase *entities.Purchase) error {
	ingredient, err := s.repo.FindByID(ctx, purchase.IngredientID)
	if err != nil {
		return err
	}
	if ingredient == nil {
		return errors.New("ingredient not found")
	}

	ingredient.CurrentStock += purchase.Quantity
	ingredient.CostPerUnit = purchase.UnitCost
	ingredient.UpdatedAt = time.Now()

	return s.repo.Update(ctx, purchase.IngredientID, ingredient)
}

// GetStockLevels returns stock levels for all ingredients
func (s *IngredientService) GetStockLevels(ctx context.Context) ([]*entities.Ingredient, error) {
	_, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	// TODO:
	return nil, nil
}