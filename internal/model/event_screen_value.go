package model

type eventValue struct {
	Column int
}

func (i eventValue) DrawHeader(index int, e Event, s *ScreenTD) {
	s.Set(index, i.GetColumn(), "Value")
}

func (i eventValue) DrawLine(index int, e Event, s *ScreenTD) {
	s.Set(index, i.GetColumn(), e.Value)
}

func (i eventValue) DrawFooter(index int, e Event, s *ScreenTD) {
	s.Set(index, i.GetColumn(), e.Value)
}

func (i eventValue) GetColumn() int {
	return i.Column
}

func NewFieldEventValue(column int) EventScreenField {
	return &eventValue{
		Column: column,
	}
}
