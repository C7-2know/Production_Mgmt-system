package interfaces

import (
	"context"
	"time"

	"soap_factory/domain/entities"
)

type ProductionService interface {
	CreateBatch(ctx context.Context, batch *entities.ProductionBatch) error
	StartBatch(ctx context.Context, batchID string) error
	CompleteBatch(ctx context.Context, batchID string, actualCount int, actualUsage []entities.IngredientUsage) error
	GetBatch(ctx context.Context, batchID string) (*entities.ProductionBatch, error)
	ListBatches(ctx context.Context, filter BatchFilter) ([]*entities.ProductionBatch, error)
	CalculateWastage(ctx context.Context, batchID string) (*entities.WastageReport, error)
}

type BatchFilter struct {
	Status    entities.BatchStatus
	StartDate time.Time
	EndDate   time.Time
	FormulaID string
}