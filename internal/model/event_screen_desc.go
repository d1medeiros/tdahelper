package model

type eventDesc struct {
	Column int
}

func (i eventDesc) DrawHeader(index int, e Event, s *ScreenTD) {
	s.Set(index, i.GetColumn(), "Desc")
}

func (i eventDesc) DrawLine(index int, e Event, s *ScreenTD) {
	s.Set(index, i.GetColumn(), e.Desc)
}

func (i eventDesc) DrawFooter(index int, e Event, s *ScreenTD) {
	s.Set(index, i.GetColumn(), "")
}

func (i eventDesc) GetColumn() int {
	return i.Column
}

func NewFieldEventDesc(column int) EventScreenField {
	return &eventDesc{
		Column: column,
	}
}
