package model

import "time"

type eventDuoDate struct {
	Column int
}

func (i eventDuoDate) DrawHeader(index int, e Event, s *ScreenTD) {
	s.Set(index, i.GetColumn(), "DuoDate")
}

func (i eventDuoDate) DrawLine(index int, e Event, s *ScreenTD) {
	s.Set(index, i.GetColumn(), e.Date.Format(time.DateOnly))
}

func (i eventDuoDate) DrawFooter(index int, e Event, s *ScreenTD) {
	s.Set(index, i.GetColumn(), "")
}

func (i eventDuoDate) GetColumn() int {
	return i.Column
}

func NewFieldEventDuoDate(column int) EventScreenField {
	return &eventDuoDate{
		Column: column,
	}
}
