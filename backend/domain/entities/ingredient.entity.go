package entities

import "time"

type Ingredient struct {
	ID           string    `bson:"_id,omitempty" json:"id"`
	Name         string    `bson:"name" json:"name"`
	Unit         string    `bson:"unit" json:"unit"` // kg, liter, piece
	CurrentStock float64   `bson:"current_stock" json:"current_stock"`
	MinStock     float64   `bson:"min_stock" json:"min_stock"`
	CostPerUnit  float64   `bson:"cost_per_unit" json:"cost_per_unit"`
	CreatedAt    time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at" json:"updated_at"`
}