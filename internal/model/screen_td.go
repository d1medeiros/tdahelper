package model

type ScreenTD struct {
	td [][]string
}

func NewScreenTD(x int, y int) *ScreenTD {
	twoDSlice1 := make([][]string, x)
	for i := range twoDSlice1 {
		twoDSlice1[i] = make([]string, y)
	}
	return &ScreenTD{td: twoDSlice1}
}

func (s *ScreenTD) GetTD() [][]string {
	return s.td
}

func (s *ScreenTD) Set(x int, y int, value string) {
	s.td[x][y] = value
}
