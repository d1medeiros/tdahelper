package model

type eventCategory struct {
	Column int
}

func (i eventCategory) DrawHeader(index int, e Event, s *ScreenTD) {
	s.Set(index, i.GetColumn(), "Category")
}

func (i eventCategory) DrawLine(index int, e Event, s *ScreenTD) {
	s.Set(index, i.GetColumn(), e.Category)
}

func (i eventCategory) DrawFooter(index int, e Event, s *ScreenTD) {
	s.Set(index, i.GetColumn(), "")
}

func (i eventCategory) GetColumn() int {
	return i.Column
}

func NewFieldEventCategory(column int) EventScreenField {
	return &eventCategory{
		Column: column,
	}
}
