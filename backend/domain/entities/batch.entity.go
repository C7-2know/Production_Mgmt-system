type Batch struct {
	ID          string  `json:"id"`
	SoapIDs     []string  `json:"soap_ids"`
	Quantity    int     `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
