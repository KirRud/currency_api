package database

import (
	"currency_api/internal/models"
	"time"
)

type CurrencyDataBaseInterface interface {
	GetRateByCurrencyAndDate(date time.Time, currency string) (models.CurrencyDB, error)
	UpdateCurrencyRate(sqlStr string, values []interface{}) error
}

func (d *DataBase) GetRateByCurrencyAndDate(date time.Time, currency string) (models.CurrencyDB, error) {
	var result models.CurrencyDB
	if err := d.db.Raw("SELECT * FROM currency_dbs WHERE date = ? AND char_code = ?", date.Format("02/01/2006"), currency).
		Scan(&result).Error; err != nil {
		return models.CurrencyDB{}, err
	}
	return result, nil
}

func (d *DataBase) UpdateCurrencyRate(sqlStr string, values []interface{}) error {
	if err := d.db.Exec(sqlStr, values...).Error; err != nil {
		return err
	}
	return nil
}
