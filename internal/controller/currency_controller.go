package controller

import (
	"currency_api/internal/database"
	"currency_api/internal/models"
	"currency_api/internal/usecase"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type CurrencyControllerInterface interface {
	GetCurrencyRate(date time.Time, currency string) (string, error)
	updateCurrency(date time.Time) error
	RunScheduler()
}

type CurrencyController struct {
	db      database.DatabaseInterface
	usecase usecase.CurrencyUseCaseInterface
}

func NewCurrencyController(db database.DatabaseInterface, uc usecase.CurrencyUseCaseInterface) CurrencyControllerInterface {
	return CurrencyController{db, uc}
}

func (cc CurrencyController) GetCurrencyRate(date time.Time, currency string) (string, error) {
	if date.After(time.Now()) {
		return "", fmt.Errorf("error date after today")
	}
	currencyRateDB, err := cc.db.GetRateByCurrencyAndDate(date, currency)
	if err != gorm.ErrRecordNotFound && err != nil {
		return "", err
	}

	if currencyRateDB != (models.CurrencyDB{}) {
		return currencyRateDB.VunitRate, nil
	}
	currencyRateList, err := cc.usecase.GetAllCurrenciesRate(date)
	if err != nil {
		return "", err
	}

	var result string
	for _, currencyRate := range currencyRateList {
		if currencyRate.CharCode == currency {
			result = currencyRate.VunitRate
			break
		}
	}
	if result == "" {
		return "", fmt.Errorf("error not found currency rate")
	}
	return result, nil
}

func (cc CurrencyController) updateCurrency(date time.Time) error {
	currencyRateList, err := cc.usecase.GetAllCurrenciesRate(date)
	if err != nil {
		return err
	}

	var sb strings.Builder
	var valuesSql []interface{}
	sb.WriteString("INSERT OR REPLACE INTO currency_dbs(id, num_code, char_code, nominal, name, value, vunit_rate, date) VALUES ")
	for i, currencyRate := range currencyRateList {
		sb.WriteString("(?,?,?,?,?,?,?,?)")
		if i != len(currencyRateList)-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(";")
		}
		valuesSql = append(valuesSql, currencyRate.ID, currencyRate.NumCode, currencyRate.CharCode, currencyRate.Nominal,
			currencyRate.Name, currencyRate.Value, currencyRate.VunitRate, date.Format("02/01/2006"))
	}

	return cc.db.UpdateCurrencyRate(sb.String(), valuesSql)
}
