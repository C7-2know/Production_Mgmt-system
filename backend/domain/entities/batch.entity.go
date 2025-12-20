package entities

import(
	"time"
)

type ProductionBatch struct {
	ID              string    `bson:"_id,omitempty" json:"id"`
	FormulaID       string    `bson:"formula_id" json:"formula_id"`
	BatchNumber     string    `bson:"batch_number" json:"batch_number"`
	ProductionDate  time.Time `bson:"production_date" json:"production_date"`
	
	// Planned quantities
	PlannedSoapCount int       `bson:"planned_soap_count" json:"planned_soap_count"`
	PlannedWeight    float64   `bson:"planned_weight" json:"planned_weight"`
	
	// Actual results
	ActualSoapCount  int       `bson:"actual_soap_count" json:"actual_soap_count"`
	ActualWeight     float64   `bson:"actual_weight" json:"actual_weight"`
	
	// Ingredient usage
	IngredientUsage []IngredientUsage `bson:"ingredient_usage" json:"ingredient_usage"`
	
	Status          BatchStatus `bson:"status" json:"status"` // planned, in-progress, completed
	Notes           string      `bson:"notes" json:"notes"`
	CreatedAt       time.Time   `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time   `bson:"updated_at" json:"updated_at"`
}

type IngredientUsage struct {
	IngredientID string  `bson:"ingredient_id" json:"ingredient_id"`
	PlannedQty   float64 `bson:"planned_qty" json:"planned_qty"`
	ActualQty    float64 `bson:"actual_qty" json:"actual_qty"`
	Wastage      float64 `bson:"wastage" json:"wastage"`
}

type BatchStatus string

const (
	BatchStatusPlanned     BatchStatus = "planned"
	BatchStatusInProgress  BatchStatus = "in_progress"
	BatchStatusCompleted   BatchStatus = "completed"
	BatchStatusCancelled   BatchStatus = "cancelled"
)