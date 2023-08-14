package model

import (
	"github.com/jameycribbs/hare"
	"sort"
	"strconv"
	"time"
)

func SortEventsById(events []Event) {
	sort.Sort(EventById(events))
}

func SortEventsByDate(events []Event) {
	sort.Sort(EventByDate(events))
}

type EventById []Event

func (a EventById) Len() int           { return len(a) }
func (a EventById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a EventById) Less(i, j int) bool { return a[i].Id < a[j].Id }

type EventByDate []Event

func (a EventByDate) Len() int           { return len(a) }
func (a EventByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a EventByDate) Less(i, j int) bool { return a[i].Date.Before(a[j].Date) }

type Event struct {
	Id          int       `json:"id"`
	Desc        string    `json:"desc"`
	Value       string    `json:"value"`
	Date        time.Time `json:"date"`
	Category    string    `json:"category"`
	Type        string    `json:"type"`
	Father      int       `json:"father"`
	Installment int       `json:"installment"`
	InsQtd      int       `json:"ins_qtd"`
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
