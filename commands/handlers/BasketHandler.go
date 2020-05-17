package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"scs/entities"
	"scs/interfaces"
	"time"
)

type Message struct {
	Message string
}

var (
	ErrorBasketNotFound         = Message{Message: "Basket not found"}
	ErrorBasketItemNotFound     = Message{Message: "Basket item not found"}
	ErrorBasketItemBadSignature = Message{Message: "Basket item bad sgnature"}
	BasketSuccessfulCreated     = Message{Message: "Create basket operation success"}
	BasketSuccessfulAppend      = Message{Message: "Append basket item operation success"}
	BasketSuccessfulUpdate      = Message{Message: "Update basket item operation success"}
	BasketSuccessfulRemove      = Message{Message: "Remove basket operation success"}
)

type BasketHandler struct {
	storage         interfaces.TemporaryStorageInterface
	storageDuration time.Duration
}

func Init(storage interfaces.TemporaryStorageInterface) BasketHandler {
	return BasketHandler{storage: storage}
}
func (bh *BasketHandler) getBasket(key string) (entities.Basket, bool) {
	if basket, exists := bh.storage.Get(key); exists == true {
		return basket.(entities.Basket), true
	}
	return entities.Basket{}, false
}

func (bh *BasketHandler) setBasket(basket entities.Basket) {
	bh.storage.Set(basket.Id.String(), basket, bh.storageDuration)
}

func (bh *BasketHandler) removeBasket(basket entities.Basket) {
	if exists == false {
		return context.JSON(http.StatusNotFound, ErrorBasketNotFound)
	}
	return context.JSON(http.StatusOK, basket)
}

func (bh *BasketHandler) CreateBasket(context echo.Context) error {
	basket := entities.CreateBasket()
	bh.setBasket(basket)
	return context.JSON(http.StatusCreated, BasketSuccessfulCreated)
}

func (bh BasketHandler) GetBasket(context echo.Context) error {
	basket, exists := bh.getBasket(context.Param("basketId"))
	if exists == false {
		return context.JSON(http.StatusNotFound, ErrorBasketNotFound)
	}
	return context.JSON(http.StatusOK, basket)
}

func (bh *BasketHandler) PostItemBasket(c echo.Context) error {
	basket, exists := bh.getBasket(c.Param("basketId"))
	if exists == false {
		return c.JSON(http.StatusNotFound, ErrorBasketNotFound)
	}
	var basketItem entities.BasketItem
	if basketItemError := c.Bind(&basketItem); basketItemError != nil {
		return c.JSON(http.StatusBadRequest, ErrorBasketItemBadSignature)
	}

	basket.Add(basketItem)
	bh.setBasket(basket)
	return c.JSON(http.StatusCreated, BasketSuccessfulAppend)
}

func (bh *BasketHandler) PutItemBasket(c echo.Context) error {
	basket, exists := bh.getBasket(c.Param("basketId"))
	if exists == false {
		return c.JSON(http.StatusNotFound, ErrorBasketNotFound)
	}

	itemId := c.Param("itemId")
	basketItem, basketItemExists := basket.Items[itemId]
	if basketItemExists == false {
		return c.JSON(http.StatusNotFound, ErrorBasketItemNotFound)
	}

	if basketItemError := c.Bind(&basketItem); basketItemError != nil {
		return c.JSON(http.StatusBadRequest, ErrorBasketItemBadSignature)
	}

	basket.Update(basketItem)
	bh.setBasket(basket)
	return c.JSON(http.StatusAccepted, BasketSuccessfulUpdate)
}

func (bh *BasketHandler) Flush(c echo.Context) error {
	basket, exists := bh.getBasket(c.Param("basketId"))
	if exists == false {
		return c.JSON(http.StatusNotFound, ErrorBasketNotFound)
	}

	bh.removeBasket(basket)
	return c.JSON(http.StatusAccepted, BasketSuccessfulRemove)
}
