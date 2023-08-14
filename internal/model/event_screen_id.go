package model

import "fmt"

type eventId struct {
	Column int
}

func (i eventId) DrawHeader(index int, e Event, s *ScreenTD) {
	s.Set(index, i.GetColumn(), "Id")
}

func (i eventId) DrawLine(index int, e Event, s *ScreenTD) {
	s.Set(index, i.GetColumn(), fmt.Sprintf("%d", e.Id))
}

func (i eventId) DrawFooter(index int, e Event, s *ScreenTD) {
	s.Set(index, i.GetColumn(), "")
}

func (i eventId) GetColumn() int {
	return i.Column
}

func NewFieldEventId(column int) EventScreenField {
	return &eventId{
		Column: column,
	}
}
