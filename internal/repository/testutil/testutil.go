package testutil

import (
	"fmt"
	"os/exec"
)

func InitDatabaseTest(path string, fileCP string, fileDB string) {
	fmt.Printf("Disk test starting")
	f := fmt.Sprintf("%s/%s.txt", path, fileCP)
	fmt.Printf("cp file: %s to %s\n", f, fileDB)
	cmd := exec.Command("cp", f, fileDB)
	if err := cmd.Run(); err != nil {
		panic(err)
	}

}
