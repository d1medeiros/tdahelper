package config

import (
	"github.com/jameycribbs/hare"
	"github.com/jameycribbs/hare/datastores/disk"
)

type DBStore struct {
	db *hare.Database
	ds *disk.Disk
}

func (dbs *DBStore) Close() {
	_ = dbs.db.Close()
}

func (dbs *DBStore) GetDB() *hare.Database {
	return dbs.db
}

func (dbs *DBStore) New(path string) {

	ds, err := disk.New(path, ".json")
	if err != nil {
		panic(err)
	}
	//fmt.Println("Disk started")
	dbs.ds = ds

	//fmt.Println("Ram starting")
	db, err := hare.New(ds)
	if err != nil {
		panic(err)
	}
	//fmt.Println("Ram started")
	dbs.db = db
}
