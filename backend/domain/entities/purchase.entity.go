package entities

import "time"

type Purchase struct {
	ID           string    `bson:"_id,omitempty" json:"id"`
	IngredientID string    `bson:"ingredient_id" json:"ingredient_id"`
	Quantity     float64   `bson:"quantity" json:"quantity"`
	UnitCost     float64   `bson:"unit_cost" json:"unit_cost"`
	TotalCost    float64   `bson:"total_cost" json:"total_cost"`
	PurchaseDate time.Time `bson:"purchase_date" json:"purchase_date"`
	Supplier     string    `bson:"supplier" json:"supplier"`
	CreatedAt    time.Time `bson:"created_at" json:"created_at"`
}