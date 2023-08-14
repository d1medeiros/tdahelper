package subop

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"tdahelper/config"
	"time"
)

var OpBKPCmd = &cobra.Command{
	Use: "bkp",
	Run: func(cmd *cobra.Command, args []string) {
		bkp(config.Path, "events.json")
	},
}

func bkp(path string, fileDB string) {
	name := time.Now().Format(time.RFC3339)
	fileCP := fmt.Sprintf("%s/%s.bkp", path, name)
	fileDB = fmt.Sprintf("%s/%s", path, fileDB)
	fmt.Printf("backup for file: %s\n", fileCP)
	cmd := exec.Command("cp", fileDB, fileCP)
	if err := cmd.Run(); err != nil {
		panic(err)
	}

}
