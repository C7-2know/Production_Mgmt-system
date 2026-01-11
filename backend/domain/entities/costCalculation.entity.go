package entities

import "time"

type SoapCost struct {
	BatchID         string  `bson:"batch_id" json:"batch_id"`
	SoapCount       int     `bson:"soap_count" json:"soap_count"`
	
	// Costs
	IngredientCost  float64 `bson:"ingredient_cost" json:"ingredient_cost"`
	PackagingCost   float64 `bson:"packaging_cost" json:"packaging_cost"`
	LaborCost       float64 `bson:"labor_cost" json:"labor_cost"`
	TotalCost       float64 `bson:"total_cost" json:"total_cost"`
	
	// Per soap
	CostPerSoap     float64 `bson:"cost_per_soap" json:"cost_per_soap"`
	
	CalculatedAt    time.Time `bson:"calculated_at" json:"calculated_at"`
}