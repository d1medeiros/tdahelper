package model

import (
	"github.com/jameycribbs/hare"
	"strconv"
	"time"
)

type Event struct {
	Id       int       `json:"id"`
	Desc     string    `json:"desc"`
	Value    string    `json:"value"`
	Date     time.Time `json:"date"`
	Category string    `json:"category"`
}

func (e *Event) SetID(i int) {
	e.Id = i
}

func (e *Event) GetID() int {
	return e.Id
}

func (e *Event) GetValueFloat64() float64 {
	float, err := strconv.ParseFloat(e.Value, 64)
	if err != nil {
		return 0
	}
	return float
}

func (e *Event) AfterFind(database *hare.Database) error {
	*e = Event(*e)
	return nil
}
