package interfaces

import (
	"context"
	"time"

	"soap_factory/domain/entities"
)

type ReportService interface {
	GenerateWeeklyReport(ctx context.Context, weekStart time.Time) (*entities.WeeklyReport, error)
	GenerateMonthlyReport(ctx context.Context, month time.Time) (*entities.MonthlyReport, error)
	GenerateWastageReport(ctx context.Context, filter ReportFilter) (*entities.WastageReport, error)
	GenerateStockReport(ctx context.Context) (*entities.StockReport, error)
}

type ReportFilter struct {
	StartDate time.Time
	EndDate   time.Time
}