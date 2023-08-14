package model

import "fmt"

type eventType struct {
	Column int
}

func (i eventType) DrawHeader(index int, e Event, s *ScreenTD) {
	s.Set(index, i.GetColumn(), "Type")
}

func (i eventType) DrawLine(index int, e Event, s *ScreenTD) {
	t := e.Type
	if e.InsQtd > 0 {
		t += fmt.Sprintf("(%d/%d)", e.Installment, e.InsQtd)
	}
	s.Set(index, i.GetColumn(), t)
}

func (i eventType) DrawFooter(index int, e Event, s *ScreenTD) {
	s.Set(index, i.GetColumn(), "")
}

func (i eventType) GetColumn() int {
	return i.Column
}

func NewFieldEventType(column int) EventScreenField {
	return &eventType{
		Column: column,
	}
}
