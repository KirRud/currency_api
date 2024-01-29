package usecase

import (
	"currency_api/internal/models"
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"net/http"
	"strings"
	"time"
)

const baseURL = "http://www.cbr.ru/scripts/XML_daily.asp"

//const baseURL = "https://www.cbr-xml-daily.ru/daily.xml"

type CurrencyUseCaseInterface interface {
	GetAllCurrenciesRate(date time.Time) ([]models.Currency, error)
}

type CurrencyUseCase struct {
	client *http.Client
}

func NewCurrencyUseCase(client *http.Client) CurrencyUseCaseInterface {
	return CurrencyUseCase{client}
}

type xmlResult struct {
	ValCurs  xml.Name          `xml:"ValCurs"`
	Date     string            `xml:"Date,attr"`
	Currency []models.Currency `xml:"Valute"`
}

func (cu CurrencyUseCase) GetAllCurrenciesRate(date time.Time) ([]models.Currency, error) {
	url := getDailyURL(date)
	resp, err := cu.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data xmlResult
	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	return data.Currency, nil

}

func getDailyURL(date time.Time) string {
	var sb strings.Builder
	sb.WriteString(baseURL)
	sb.WriteString("?date_req=")
	sb.WriteString(date.Format("02/01/2006"))
	return sb.String()
}
