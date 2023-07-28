package bulkac

import (
	"github.com/spf13/cobra"
	"tdahelper/cmd/bulkac/bulksubac"
)

func init() {
	BulkCmd.AddCommand(bulksubac.EventsCmd)
}

var BulkCmd = &cobra.Command{
	Use: "bulk",
}
