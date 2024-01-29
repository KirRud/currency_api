package main

import (
	"currency_api/config"
	"currency_api/internal/controller"
	"currency_api/internal/database"
	"currency_api/internal/handler"
	"currency_api/internal/usecase"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal("Error loading configuration: ", err)
		return
	}

	db, err := database.InitDB(cfg.DB)
	if err != nil {
		log.Fatal("Error initializing database:", err)
		return
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	uc := usecase.NewCurrencyUseCase(client)
	c := controller.NewCurrencyController(db, uc)
	h := handler.NewHandler(c)

	go c.RunScheduler()
	router := handler.InitRoutes(h)
	router.Run(":" + cfg.Server.Port)

}
