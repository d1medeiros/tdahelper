package config

import (
	"errors"
	"fmt"
	"github.com/artonge/go-csv-tag/v2"
	"tdahelper/internal/model"
	"tdahelper/internal/repository"
	"time"
)

type Manager struct {
	er repository.EventRepository
}

func NewManager(er repository.EventRepository) Manager {
	return Manager{
		er: er,
	}
}

func findWrapper(all []model.Event) [][]string {
	td := make([][]string, 0)
	line := []string{"Desc", "Value", "Category"}
	td = append(td, line)
	for _, ev := range all {
		line := []string{ev.Desc, ev.Value, ev.Category}
		td = append(td, line)
	}
	return td
}

func (m Manager) FindByMouth(mouth int, year int) [][]string {
	all, err := m.er.FindBy(
		func(ev model.Event) bool {
			return ev.Date.Month() == time.Month(mouth) && ev.Date.Year() == year
		},
	)
	if err != nil {
		panic(err)
	}
	return findWrapper(all)
}

func (m Manager) Create(desc string, value string, date time.Time, category string) error {
	e := model.Event{
		Desc:     desc,
		Value:    value,
		Date:     date,
		Category: category,
	}
	insert, err := m.er.Insert(e)
	if err != nil {
		return err
	}
	if insert != 1 {
		return errors.New("was not created")
	}
	return nil
}

func (m Manager) FindAll() [][]string {
	all, err := m.er.FindAll()
	if err != nil {
		panic(err)
	}
	return findWrapper(all)
}

func (m Manager) Bulk(fullFilePath string) error {
	var tab []model.EventFileItem
	err := csvtag.LoadFromPath(
		fullFilePath,
		&tab,
		csvtag.CsvOptions{
			Separator: ';',
		})
	if err != nil {
		return err
	}
	var errs = make([]string, 0)
	for i, item := range tab {
		e := item.ToEvent()
		_, err := m.er.Insert(e)
		if err != nil {
			errs = append(
				errs,
				fmt.Sprintf(
					"line:%d desc:%s date:%s %s",
					i,
					item.Desc,
					item.Date,
					err,
				),
			)
		}
	}

	if len(errs) >= 1 {
		for _, s := range errs {
			fmt.Println(s)
		}
	}

	return nil
}
