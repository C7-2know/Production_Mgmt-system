package entities

import "time"

type Packaging struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Name        PackagingMaterial    `bson:"name" json:"name"` // plastic wrap, label, box
	Type        string    `bson:"type" json:"type"` // wrap, label, box
	Unit        string    `bson:"unit" json:"unit"` // piece, roll, sheet
	CostPerUnit float64   `bson:"cost_per_unit" json:"cost_per_unit"`
	CurrentStock int      `bson:"current_stock" json:"current_stock"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}

type PackagingUsage struct {
	PackagingID string `bson:"packaging_id" json:"packaging_id"`
	Quantity    int    `bson:"quantity" json:"quantity"`
}

type PackagingMaterial struct {
	PlasticWrap string
	Label       string
	Box         string
}