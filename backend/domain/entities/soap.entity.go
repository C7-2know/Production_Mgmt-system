type Soap struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
	Ingredients []string `json:"ingredients"`
	BatchIDs  []string `json:"batch_ids"`
}