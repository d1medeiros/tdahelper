package listac

import (
	"github.com/spf13/cobra"
	"tdahelper/cmd/listac/listasubac"
)

func init() {
	ListEventsCmd.AddCommand(listasubac.EventsCmd)
}

var ListEventsCmd = &cobra.Command{
	Use:   "list",
	Short: "Print the * based on filter",
}
