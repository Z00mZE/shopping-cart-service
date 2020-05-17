package entities

import (
	"github.com/google/uuid"
	"time"
)

type Basket struct {
	Id        uuid.UUID
	CreatedAt time.Time
	Items     map[string]BasketItem
}

func (basket *Basket) GetItems() map[string]BasketItem {
	return basket.Items
}
func CreateBasket() Basket {
	basket := Basket{
		Id:        uuid.New(),
		CreatedAt: time.Now(),
		Items:     make(map[string]BasketItem),
	}

	return basket
}
func (basket *Basket) Add(item BasketItem) {
	item.Id = uuid.New()
	item.CreatedAt = time.Now()
	basket.setItem(item)
}
func (basket *Basket) Update(item BasketItem) {
	basket.setItem(item)
}
func (basket *Basket) setItem(item BasketItem) {
	item.UpdateAt = time.Now()
	basket.Items[item.Id.String()] = item
}
