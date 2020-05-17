package main

import (
	"github.com/akyoto/cache"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
	"scs/handlers"
	"time"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	cacheStorage := cache.New(24 * time.Hour)
	BasketHandler := handlers.Init(cacheStorage)
	// Routes
	e.GET("/create", BasketHandler.CreateBasket)

	g := e.Group("/basket/:basketId")
	{
		g.GET("", BasketHandler.GetBasket)
		g.DELETE("",BasketHandler.Flush)
		g.POST("", BasketHandler.PostItemBasket)
		g.PUT("/:itemId", BasketHandler.PutItemBasket)
	}

	s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}
	// Start server
	e.Logger.Fatal(e.StartH2CServer(":80", s))
}
