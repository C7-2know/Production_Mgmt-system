package interfaces

import (
	"context"

	"soap_factory/domain/entities"
)

type FormulaService interface {
	CreateFormula(ctx context.Context, formula *entities.SoapFormula) error
	UpdateFormula(ctx context.Context, id string, formula *entities.SoapFormula) error
	GetFormula(ctx context.Context, id string) (*entities.SoapFormula, error)
	ListFormulas(ctx context.Context) ([]*entities.SoapFormula, error)
	CalculateBatchRequirements(ctx context.Context, formulaID string, soapCount int) (BatchRequirements, error)
}

type BatchRequirements struct {
	Formula      *entities.SoapFormula
	SoapCount    int
	Ingredients  []IngredientRequirement
	TotalWeight  float64
}

type IngredientRequirement struct {
	Ingredient *entities.Ingredient
	Quantity   float64
}