package main

import (
	"fmt"
	"github.com/jameycribbs/hare"
	"github.com/jameycribbs/hare/datastores/disk"
	"tdahelper/internal/model"
	"tdahelper/internal/repository"
	"time"
)

func main() {
	fmt.Println("Disk starting")
	ds, err := disk.New("./data", ".json")
	if err != nil {
		panic(err)
	}
	fmt.Println("Disk started")
	fmt.Println("Ram starting")
	db, err := hare.New(ds)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Ram started")
	er := repository.NewEventRepository(db)

	e := &model.Event{
		Id:       2,
		Desc:     "Ifood",
		Value:    "10.05",
		Date:     time.Now(),
		Category: "Restaurante",
	}

	_, err = er.Insert(e)
	if err != nil {
		panic(err)
	}

	events, err := er.FindAll()
	if err != nil {
		panic(err)
	}
	for _, i := range events {
		fmt.Println(i.Desc)
		fmt.Println(i.Value)
		fmt.Println("--------")
	}

}
