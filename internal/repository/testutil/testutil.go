package testutil

import (
	"fmt"
	"github.com/jameycribbs/hare"
	"github.com/jameycribbs/hare/datastores/disk"
	"os/exec"
)

func InitDatabaseTest(path string, fileCP string, fileDB string) *hare.Database {
	fmt.Println("Disk test starting")
	f := fmt.Sprintf("%s/%s.txt", path, fileCP)
	fmt.Printf("cp file: %s to %s", f, fileDB)
	cmd := exec.Command("cp", f, fileDB)
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	ds, err := disk.New(path, ".json")
	if err != nil {
		panic(err)
	}
	fmt.Println("Disk test started")

	fmt.Println("Ram test starting")
	db, err := hare.New(ds)
	if err != nil {
		panic(err)
	}
	fmt.Println("Ram test started")

	return db
}
