package entities

import (
	"github.com/google/uuid"
	"time"
)

type BasketItems struct {
	items map[string]BasketItem
}

func InitBasketItemsCollection() BasketItems {
	return BasketItems{items: make(map[string]BasketItem)}
}

func (bi *BasketItems) Flush() {
	bi.items = make(map[string]BasketItem)
}

/** Read */
func (bi *BasketItems) GetItem(key string) BasketItem {
	return bi.items[key]
}

/** Create */
func (bi *BasketItems) Add(item BasketItem) {
	item.Id = uuid.New()
	item.CreatedAt = time.Now()
	item.UpdateAt = time.Now()
	bi.items[item.Id.String()] = item
}

/** Update */
func (bi *BasketItems) Update(item BasketItem) {
	bi.items[item.Id.String()] = item
}

/** Delete */
func (bi *BasketItems) Delete(item BasketItem) {
	delete(bi.items, item.Id.String())
}
