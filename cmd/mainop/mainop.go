package mainop

import (
	"github.com/spf13/cobra"
	"tdahelper/cmd/mainop/subop"
)

func init() {
	OpCmd.AddCommand(subop.OpCPCmd)
	OpCmd.AddCommand(subop.OpCleanCmd)
	OpCmd.AddCommand(subop.OpBKPCmd)
}

var OpCmd = &cobra.Command{
	Use: "op",
}
