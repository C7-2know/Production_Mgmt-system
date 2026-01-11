package interfaces

import (
	"context"
	"time"

	"soap_factory/domain/entities"
)

type CostService interface {
	CalculateBatchCost(ctx context.Context, batchID string) (*entities.SoapCost, error)
	GetCostPerSoap(ctx context.Context, batchID string) (float64, error)
	GenerateCostReport(ctx context.Context, filter CostReportFilter) (*entities.CostReport, error)
}

type CostReportFilter struct {
	StartDate time.Time
	EndDate   time.Time
	FormulaID string
}