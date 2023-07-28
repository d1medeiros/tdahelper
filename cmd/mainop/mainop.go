package mainop

import (
	"github.com/spf13/cobra"
	"tdahelper/cmd/mainop/subop"
)

func init() {
	OpCmd.AddCommand(subop.OpCPCmd)
	OpCmd.AddCommand(subop.OpCleanCmd)
}

var OpCmd = &cobra.Command{
	Use: "op",
}
