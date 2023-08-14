package repository

import (
	"github.com/stretchr/testify/suite"
	"tdahelper/internal/model"
	"tdahelper/internal/repository/testutil"
	"testing"
	"time"
)

var path = "./datatest"
var baseDB = "./datatest/events.json"

type EventRepositoryTest struct {
	suite.Suite
	dbs *model.DBStore
}

func TestEventRepositoryTest(t *testing.T) {
	suite.Run(t, new(EventRepositoryTest))
}

func (s *EventRepositoryTest) SetupTest() {
	s.dbs = &model.DBStore{}
}

func (s *EventRepositoryTest) TearDownAllSuite() {
	s.dbs.Close()
}

func (s *EventRepositoryTest) TestEventRepositoryFindAll() {
	tests := []struct {
		name     string
		initFile string
		want     []model.Event
		wantErr  bool
	}{
		{
			name:     "findAll with success",
			initFile: "baseline-test",
			want: []model.Event{
				{
					Id:       1,
					Desc:     "teste 1",
					Value:    "2.0",
					Date:     time.Time{},
					Category: "a",
				},
				{
					Id:       2,
					Desc:     "teste 2",
					Value:    "3.0",
					Date:     time.Time{},
					Category: "a",
				},
				{
					Id:       3,
					Desc:     "teste 3",
					Value:    "3.0",
					Date:     time.Time{},
					Category: "b",
				},
				{
					Id:       4,
					Desc:     "teste 4",
					Value:    "3.0",
					Date:     time.Time{},
					Category: "b",
				},
				{
					Id:       5,
					Desc:     "teste 5",
					Value:    "3.0",
					Date:     time.Time{},
					Category: "c",
				},
			},
			wantErr: false,
		},
		{
			name:     "findAll with empty",
			initFile: "empty-test",
			want:     nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			testutil.InitDatabaseTest(path, tt.initFile, baseDB)
			s.dbs.New(path)
			e := &eventRepository{
				db: s.dbs.GetDB(),
			}
			got, err := e.FindAll()
			if (err != nil) != tt.wantErr {
				s.Errorf(err, "FindAll() wantErr %v", tt.wantErr)
				return
			}
			model.SortEventsById(got)
			s.EqualValues(
				tt.want,
				got,
			)
		})
	}
}

func (s *EventRepositoryTest) TestEventRepositoryInsert() {
	type args struct {
		r model.Event
	}
	tests := []struct {
		name     string
		args     args
		initFile string
		want     int
		wantErr  bool
	}{
		{
			name: "insert with success",
			args: args{
				r: model.Event{
					Desc:     "teste 1",
					Value:    "2.0",
					Date:     time.Time{},
					Category: "a",
				},
			},
			initFile: "empty-test",
			want:     1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			testutil.InitDatabaseTest(path, tt.initFile, baseDB)
			s.dbs.New(path)
			e := &eventRepository{
				db: s.dbs.GetDB(),
			}
			got, err := e.Insert(tt.args.r)
			if (err != nil) != tt.wantErr {
				s.Errorf(err, "Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				s.Errorf(err, "Insert() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *EventRepositoryTest) TestEventRepositoryUpdate() {
	type args struct {
		ev model.Event
	}
	tests := []struct {
		name     string
		initFile string
		args     args
		wantErr  bool
	}{
		{
			name:     "Update, with success",
			initFile: "baseline-test",
			args: args{
				ev: model.Event{
					Id:       1,
					Desc:     "teste 1",
					Value:    "2.0",
					Date:     time.Time{},
					Category: "z",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			testutil.InitDatabaseTest(path, tt.initFile, baseDB)
			s.dbs.New(path)
			e := &eventRepository{
				db: s.dbs.GetDB(),
			}
			err := e.Update(tt.args.ev)
			if (err != nil) != tt.wantErr {
				s.Errorf(err, "Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func (s *EventRepositoryTest) TestEventRepository_findBy() {
	type args struct {
		queryFn func(ev model.Event) bool
	}
	tests := []struct {
		name     string
		args     args
		initFile string
		want     []model.Event
		wantErr  bool
	}{
		{
			name: "FindBy [desc] with success",
			args: args{
				queryFn: func(ev model.Event) bool {
					return ev.Desc == "teste 1"
				},
			},
			initFile: "findby-equal-desc-success",
			want: []model.Event{
				{
					Id:       1,
					Desc:     "teste 1",
					Value:    "2.0",
					Date:     time.Time{},
					Category: "a",
				},
				{
					Id:       3,
					Desc:     "teste 1",
					Value:    "3.0",
					Date:     time.Time{},
					Category: "a",
				},
				{
					Id:       5,
					Desc:     "teste 1",
					Value:    "3.0",
					Date:     time.Time{},
					Category: "a",
				},
			},
			wantErr: false,
		},
		{
			name: "FindBy [category] with success",
			args: args{
				queryFn: func(ev model.Event) bool {
					return ev.Category == "b"
				},
			},
			initFile: "findby-equal-desc-success",
			want: []model.Event{
				{
					Id:       4,
					Desc:     "teste 4",
					Value:    "3.0",
					Date:     time.Time{},
					Category: "b",
				},
			},
			wantErr: false,
		},
		{
			name: "FindBy [range date] with success",
			args: args{
				queryFn: func(ev model.Event) bool {
					return ev.Date.After(time.Date(
						2023,
						2,
						1,
						0,
						0,
						0,
						0,
						time.UTC,
					)) &&
						ev.Date.Before(time.Date(
							2023,
							3,
							1,
							0,
							0,
							0,
							0,
							time.UTC,
						))
				},
			},
			initFile: "findby-range-date-success",
			want: []model.Event{
				{
					Id:    2,
					Desc:  "teste 2",
					Value: "3.0",
					Date: time.Date(
						2023,
						2,
						15,
						0,
						0,
						0,
						0,
						time.UTC,
					),
					Category: "a",
				},
				{
					Id:    3,
					Desc:  "teste 3",
					Value: "3.0",
					Date: time.Date(
						2023,
						2,
						28,
						0,
						0,
						0,
						0,
						time.UTC,
					),
					Category: "a",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			testutil.InitDatabaseTest(path, tt.initFile, baseDB)
			s.dbs.New(path)
			e := &eventRepository{
				db: s.dbs.GetDB(),
			}
			got, err := e.FindBy(tt.args.queryFn)
			if (err != nil) != tt.wantErr {
				s.Errorf(err, "FindBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			model.SortEventsById(got)
			s.EqualValues(tt.want, got)
		})
	}
}
