package subop

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"tdahelper/config"
)

var OpCPCmd = &cobra.Command{
	Use: "cp",
	Run: func(cmd *cobra.Command, args []string) {
		cp(config.Path, "events", "events.json")
	},
}

func cp(path string, fileCP string, fileDB string) {
	fileCP = fmt.Sprintf("%s/%s.txt", path, fileCP)
	fileDB = fmt.Sprintf("%s/%s", path, fileDB)
	fmt.Printf("cp file: %s to %s\n", fileCP, fileDB)
	cmd := exec.Command("cp", fileCP, fileDB)
	if err := cmd.Run(); err != nil {
		panic(err)
	}

}
