package interfaces

import (
	"context"

	"soap_factory/domain/entities"
)

type PackagingService interface {
	AddPackaging(ctx context.Context, packaging *entities.Packaging) error
	UpdatePackaging(ctx context.Context, id string, packaging *entities.Packaging) error
	RecordPackagingUsage(ctx context.Context, batchID string, usage []entities.PackagingUsage) error
	GetPackagingStock(ctx context.Context) ([]*entities.Packaging, error)
}

type PackagingRepository interface {
	InsertPackaging(ctx context.Context, packaging *entities.Packaging) error
	UpdatePackaging(ctx context.Context, id string, packaging *entities.Packaging) error
	RecordPackagingUsage(ctx context.Context, batchID string, usage []entities.PackagingUsage) error
	FetchPackagingStock(ctx context.Context) ([]*entities.Packaging, error)
}