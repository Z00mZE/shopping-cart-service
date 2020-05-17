package entities

import (
	"github.com/google/uuid"
	"time"
)

type BasketItem struct {
	Id        uuid.UUID
	CreatedAt time.Time
	UpdateAt  time.Time
	SKU       string  `json:"sku"`
	Title     string  `json:"title"`
	Quantity  uint64  `json:"quantity"`
	Price     float64 `json:"price"`
}
