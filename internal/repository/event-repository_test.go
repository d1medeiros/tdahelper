package repository

import (
	"fmt"
	"github.com/jameycribbs/hare"
	"github.com/stretchr/testify/assert"
	"tdahelper/internal/model"
	"tdahelper/internal/repository/testutil"
	"testing"
	"time"
)

var path = "./datatest"
var baseDB = "./datatest/events.json"

func initDB(db *hare.Database) {
	_ = db.DropTable(table_event)
	if !db.TableExists(table_event) {
		err := db.CreateTable(table_event)
		if err != nil {
			panic(err)
		}
	}
}

func Test_eventRepository_FindAll(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			e := &eventRepository{
				db: testutil.InitDatabaseTest(path, tt.initFile, baseDB),
			}

			got, err := e.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)

			e.db.Close()
		})
	}
}

func Test_eventRepository_Insert(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			e := &eventRepository{
				db: testutil.InitDatabaseTest(path, tt.initFile, baseDB),
			}
			got, err := e.Insert(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Insert() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eventRepository_Update(t *testing.T) {
	type args struct {
		ev model.Event
	}
	tests := []struct {
		name     string
		initFile string
		args     args
		wantErr  assert.ErrorAssertionFunc
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
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventRepository{
				db: testutil.InitDatabaseTest(path, tt.initFile, baseDB),
			}
			tt.wantErr(t, e.Update(tt.args.ev), fmt.Sprintf("Update(%v)", tt.args.ev))
			_ = e.db.Close()
		})
	}
}

func Test_eventRepository_findBy(t *testing.T) {
	type args struct {
		queryFn func(ev model.Event) bool
	}
	tests := []struct {
		name     string
		args     args
		initFile string
		want     []model.Event
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name: "findBy [desc] with success",
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
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
		{
			name: "findBy [category] with success",
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
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
		{
			name: "findBy [range date] with success",
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
						time.Now().Location(),
					)) &&
						ev.Date.Before(time.Date(
							2023,
							3,
							1,
							0,
							0,
							0,
							0,
							time.Now().Location(),
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
						time.Now().Location(),
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
						time.Now().Location(),
					),
					Category: "a",
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &eventRepository{
				db: testutil.InitDatabaseTest(path, tt.initFile, baseDB),
			}
			got, err := e.findBy(tt.args.queryFn)
			if !tt.wantErr(t, err, fmt.Sprintf("findBy()")) {
				return
			}
			assert.Equalf(t, tt.want, got, "findBy()")
			_ = e.db.Close()
		})
	}
}
