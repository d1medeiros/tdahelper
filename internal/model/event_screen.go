package model

import "fmt"

const (
	EventsById   = "SortEventsById"
	EventsByDate = "SortEventsByDate"
)

var m map[string]func(all []Event)

func init() {
	m = map[string]func(all []Event){
		EventsById: func(all []Event) {
			SortEventsById(all)
		},
		EventsByDate: func(all []Event) {
			SortEventsByDate(all)
		},
	}
}

type EventScreenField interface {
	DrawHeader(index int, e Event, s *ScreenTD)
	DrawLine(index int, e Event, s *ScreenTD)
	DrawFooter(index int, e Event, s *ScreenTD)
	GetColumn() int
}

type EventScreen struct {
	screenTD *ScreenTD
	fields   []EventScreenField
}

func NewEventScreen() *EventScreen {
	return &EventScreen{
		fields: []EventScreenField{
			NewFieldEventId(0),
			NewFieldEventDesc(1),
			NewFieldEventCategory(2),
			NewFieldEventType(3),
			NewFieldEventDuoDate(4),
			NewFieldEventValue(5),
		},
	}
}

func (es *EventScreen) Make(all []Event, ft string) [][]string {
	if _, isOk := m[ft]; isOk {
		m[ft](all)
	}
	es.screenTD = NewScreenTD(len(all)+2, len(es.fields))
	total := 0.0
	for i, e := range all {
		total += e.GetValueFloat64()
		for _, f := range es.fields {
			if i == 0 {
				f.DrawHeader(0, Event{}, es.screenTD)
			}
			f.DrawLine(i+1, e, es.screenTD)
			if i+1 == len(all) {
				eventTotal := Event{
					Value: fmt.Sprintf("%v", total),
				}
				f.DrawFooter(i+2, eventTotal, es.screenTD)
			}
		}
	}

	return es.screenTD.GetTD()
}
