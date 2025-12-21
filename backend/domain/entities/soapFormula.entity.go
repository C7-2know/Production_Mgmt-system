package entities

import "time"

type FormulaIngredient struct {
	IngredientID string  `bson:"ingredient_id" json:"ingredient_id"`
	Percentage   float64 `bson:"percentage" json:"percentage"` // 0-100%
	Weight       float64 `bson:"weight" json:"weight"`         // weight per soap
}

type SoapFormula struct {
	ID           string               `bson:"_id,omitempty" json:"id"`
	Name         string               `bson:"name" json:"name"`
	Ingredients  []FormulaIngredient  `bson:"ingredients" json:"ingredients"`
	TotalWeight  float64              `bson:"total_weight" json:"total_weight"` // per soap
	IsActive     bool                 `bson:"is_active" json:"is_active"`
	CreatedAt    time.Time            `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time            `bson:"updated_at" json:"updated_at"`
}