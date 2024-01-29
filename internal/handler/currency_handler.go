package handler

import (
	"currency_api/internal/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type HandlerInterface interface {
	GetCurrency(c *gin.Context)
}

type Handler struct {
	controller controller.CurrencyControllerInterface
}

func NewHandler(c controller.CurrencyControllerInterface) HandlerInterface {
	return Handler{c}
}

type currencyResponse struct {
	Rate     string    `json:"rate"`
	Currency string    `json:"currency"`
	Date     time.Time `json:"date"`
}

func (h Handler) GetCurrency(c *gin.Context) {
	currency := c.Query("val")
	if currency == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "empty currency"})
		return
	}
	dateStr := c.DefaultQuery("date", time.Now().Format("02.01.2006"))
	date, err := time.Parse("02.01.2006", dateStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rate, err := h.controller.GetCurrencyRate(date, currency)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := currencyResponse{rate, currency, date}
	c.IndentedJSON(http.StatusOK, response)
}
