package model

type EventFileItem struct {
	Desc     string `csv:"Desc"`
	Value    string `csv:"Value"`
	Date     string `csv:"Date"`
	Category string `csv:"Category"`
}

func (efi EventFileItem) ToEvent() Event {
	date, err := StrToDate(efi.Date)
	if err != nil {
		return Event{}
	}
	return Event{
		Desc:     efi.Desc,
		Value:    efi.Value,
		Date:     date,
		Category: efi.Category,
	}
}
