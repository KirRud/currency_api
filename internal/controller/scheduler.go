package controller

import (
	"log"
	"time"
)

func (cc CurrencyController) RunScheduler() {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Printf("Error load location: %s", err)
	}
	now := time.Now()
	nextRun := time.Date(now.Year(), now.Month(), now.Day(), 10, 0, 0, 0, loc)
	if now.After(nextRun) {
		nextRun = nextRun.AddDate(0, 0, 1)
	}
	timer := time.NewTimer(nextRun.Sub(now))
	for {
		select {
		case <-timer.C:
			now = time.Now()
			if err := cc.updateCurrency(now); err != nil {
				log.Printf("Error update currency: %s at %s", err, now.Format("02.01.2006 15:04"))
			}
			nextRun = nextRun.AddDate(0, 0, 1)
			timer.Reset(nextRun.Sub(now))
		}
	}
}
