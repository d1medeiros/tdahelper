package repository

import (
	"github.com/jameycribbs/hare"
	"tdahelper/internal/model"
)

const table_event = "events"

type EventRepository interface {
	Insert(ev model.Event) (int, error)
	Update(ev model.Event) error
	FindAll() ([]model.Event, error)
	FindByDesc(desc string) ([]model.Event, error)
}

type eventRepository struct {
	db *hare.Database
}

func NewEventRepository(db *hare.Database) EventRepository {
	if !db.TableExists(table_event) {
		err := db.CreateTable(table_event)
		if err != nil {
			panic(err)
		}
	}
	return &eventRepository{
		db: db,
	}
}

func (e *eventRepository) findBy(queryFn func(ev model.Event) bool) ([]model.Event, error) {
	var results []model.Event
	var err error
	ids, err := e.db.IDs(table_event)
	if err != nil {
		return nil, err
	}
	for _, id := range ids {
		c := model.Event{}
		if err = e.db.Find(table_event, id, &c); err != nil {
			return nil, err
		}
		if queryFn(c) {
			results = append(results, c)
		}
	}
	return results, nil
}

func (e *eventRepository) Insert(ev model.Event) (int, error) {
	return e.db.Insert(table_event, &ev)
}

func (e *eventRepository) Update(ev model.Event) error {
	return e.db.Update(table_event, &ev)
}

func (e *eventRepository) FindAll() ([]model.Event, error) {
	var results []model.Event
	var err error
	ids, err := e.db.IDs(table_event)
	if err != nil {
		return nil, err
	}
	for _, id := range ids {
		c := model.Event{}
		if err = e.db.Find(table_event, id, &c); err != nil {
			return nil, err
		}
		results = append(results, c)
	}
	return results, nil
}

func (e *eventRepository) FindByDesc(desc string) ([]model.Event, error) {
	all, err := e.findBy(
		func(ev model.Event) bool {
			return ev.Desc == desc
		},
	)
	if err != nil {
		return nil, err
	}
	return all, nil
}
