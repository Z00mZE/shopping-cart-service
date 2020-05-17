package entities

type Price struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}
