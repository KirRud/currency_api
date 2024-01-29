package controller

import (
	"currency_api/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

type MockDatabase struct{}

func (db *MockDatabase) GetRateByCurrencyAndDate(date time.Time, currency string) (models.CurrencyDB, error) {
	if currency == "USD" {
		return models.CurrencyDB{
			Currency: models.Currency{
				CharCode:  "USD",
				VunitRate: "1.2",
			},
			Date: date,
		}, nil
	}
	return models.CurrencyDB{}, gorm.ErrRecordNotFound
}

func (db *MockDatabase) UpdateCurrencyRate(sqlStr string, values []interface{}) error {
	return nil
}

type MockUseCase struct{}

func (uc *MockUseCase) GetAllCurrenciesRate(date time.Time) ([]models.Currency, error) {
	return []models.Currency{
		{
			CharCode:  "USD",
			VunitRate: "1.2",
		},
		{
			CharCode:  "EUR",
			VunitRate: "0.9",
		},
	}, nil
}

func TestGetCurrencyRate(t *testing.T) {
	db := &MockDatabase{}
	uc := &MockUseCase{}
	cc := NewCurrencyController(db, uc)

	// Test case: currency found in the database
	rate, err := cc.GetCurrencyRate(time.Now(), "USD")
	assert.Nil(t, err)
	assert.Equal(t, 1.2, rate)

	// Test case: currency not found in the database, should fetch from the use case
	rate, err = cc.GetCurrencyRate(time.Now(), "EUR")
	assert.Nil(t, err)
	assert.Equal(t, 0.9, rate)

	// Test case: date after today, should return an error
	rate, err = cc.GetCurrencyRate(time.Now().AddDate(0, 0, 1), "USD")
	assert.NotNil(t, err)
	assert.EqualError(t, err, "error date after today")

	// Test case: currency not found, should return an error
	rate, err = cc.GetCurrencyRate(time.Now(), "GBP")
	assert.NotNil(t, err)
	assert.EqualError(t, err, "error not found currency rate")
}

func TestUpdateCurrency(t *testing.T) {
	db := &MockDatabase{}
	uc := &MockUseCase{}
	cc := NewCurrencyController(db, uc)

	// Test case: update currency rates
	err := cc.updateCurrency(time.Now())
	assert.Nil(t, err)
}
