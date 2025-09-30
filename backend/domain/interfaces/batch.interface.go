package interfaces

type Batch interface {
	Fill()
	Empty()
}

type BatchRepository interface {
	Save(batch Batch) error
	FindByID(id string) (Batch, error)
	Update(batch Batch) error
	Delete(id string) error
}

type BatchService interface {
	CreateBatch(batch Batch) (Batch, error)
	GetBatchByID(id string) (Batch, error)
	UpdateBatch(batch Batch) (Batch, error)
	DeleteBatch(id string) error
}


