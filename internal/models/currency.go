package models

import (
	"time"
)

type Currency struct {
	ID        string `xml:"ID,attr" gorm:"column:id"`
	NumCode   int64  `xml:"NumCode" gorm:"column:num_code"`
	CharCode  string `xml:"CharCode" gorm:"primaryKey;column:char_code"`
	Nominal   string `xml:"Nominal" gorm:"column:nominal"`
	Name      string `xml:"Name" gorm:"column:name"`
	Value     string `xml:"Value" gorm:"column:value"`
	VunitRate string `xml:"VunitRate" gorm:"column:vunit_rate"`
}

type CurrencyDB struct {
	Currency
	Date time.Time `gorm:"column:date"`
}
