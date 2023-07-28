package subop

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"tdahelper/config"
)

var OpCleanCmd = &cobra.Command{
	Use: "clean",
	Run: func(cmd *cobra.Command, args []string) {
		clean(config.Path, "events.json")
	},
}

func clean(path string, fileDB string) {
	fileCP := fmt.Sprintf("%s/empty.txt", path)
	fileDB = fmt.Sprintf("%s/%s", path, fileDB)
	fmt.Printf("cp file: %s to %s\n", fileCP, fileDB)
	cmd := exec.Command("cp", fileCP, fileDB)
	if err := cmd.Run(); err != nil {
		panic(err)
	}

}
