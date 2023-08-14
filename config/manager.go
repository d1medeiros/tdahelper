package config

import (
	"errors"
	"fmt"
	"github.com/artonge/go-csv-tag/v2"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/spf13/pflag"
	"tdahelper/internal/model"
	"tdahelper/internal/repository"
	"time"
)

type Manager struct {
	er repository.EventRepository
	es *model.EventScreen
}

func NewManager(er repository.EventRepository) Manager {
	return Manager{
		er: er,
		es: model.NewEventScreen(),
	}
}

func (m Manager) insert(e model.Event) error {
	insert, err := m.er.Insert(e)
	if err != nil {
		return err
	}
	if insert < 1 {
		return errors.New("was not created")
	}
	return nil
}

func (m Manager) FindByMouth(mouth int, year int, fg *pflag.FlagSet) [][]string {
	sort, err := fg.GetString("sort")
	if err != nil {
		panic(err)
	}
	category, err := fg.GetString("category")
	if err != nil {
		panic(err)
	}
	all, err := m.er.FindBy(
		func(ev model.Event) bool {
			var result = ev.Date.Month() == time.Month(mouth) && ev.Date.Year() == year
			if category != "" {
				result = result && ev.Category == category
			}
			return result
		},
	)
	if err != nil {
		panic(err)
	}
	return m.es.Make(all, sort)
}

func (m Manager) Create(
	desc string,
	value string,
	date time.Time,
	category string,
	t string,
	qtd int,
) error {
	e := model.Event{
		Desc:     desc,
		Value:    value,
		Date:     date,
		Category: category,
		Type:     t,
		InsQtd:   qtd,
	}
	if qtd > 0 {
		e.Value = fmt.Sprintf("%f", e.GetValueFloat64()/float64(qtd))
		for i := 1; i <= qtd; i++ {
			if i > 1 {
				e.Father = 1
				e.Date = e.Date.AddDate(0, 1, 0)
			}
			e.Installment = i
			m.insert(e)
		}
		return nil
	}
	return m.insert(e)
}

func (m Manager) FindAll(sort string) [][]string {
	all, err := m.er.FindAll()
	if err != nil {
		panic(err)
	}
	return m.es.Make(all, sort)
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

func (m Manager) FindAllCategory() []string {
	all, err := m.er.FindAll()
	if err != nil {
		panic(err)
	}
	set := hashset.New()
	for _, e := range all {
		set.Add(e.Category)
	}
	s := make([]string, 0)
	for _, i := range set.Values() {
		s = append(s, i.(string))
	}
	return s
}

func (m Manager) FindByRange(dt1 string, dt2 string, fg *pflag.FlagSet) [][]string {
	day1, month1, year1, _ := model.StrDateSplit(dt1)
	date1 := time.Date(year1, time.Month(month1), day1, 0, 0, 0, 0, time.UTC)
	day2, month2, year2, _ := model.StrDateSplit(dt2)
	date2 := time.Date(year2, time.Month(month2), day2, 0, 0, 0, 0, time.UTC)
	sort, err := fg.GetString("sort")
	if err != nil {
		panic(err)
	}
	category, err := fg.GetString("category")
	if err != nil {
		panic(err)
	}
	all, err := m.er.FindBy(
		func(ev model.Event) bool {
			var result = ev.Date.After(date1) && ev.Date.Before(date2)
			if category != "" {
				result = result && ev.Category == category
			}
			return result
		},
	)
	if err != nil {
		panic(err)
	}
	return m.es.Make(all, sort)
}
