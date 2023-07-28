package listasubac

import (
	"github.com/spf13/cobra"
	"tdahelper/cmd/render"
	"tdahelper/config"
)

var all bool
var desc = ""
var category = ""
var byMouth = ""

var EventsCmd = &cobra.Command{
	Use:   "event",
	Short: "Print the events based on filter",
	Run: func(cmd *cobra.Command, args []string) {
		var td [][]string
		if all {
			td = config.Mng.FindAll()
		} else {
			td = config.Mng.FindByMouth(2, 2023)
		}
		_ = render.RenderTable(td)
	},
}

func init() {
	EventsCmd.Flags().BoolVarP(
		&all,
		"all",
		"a",
		false,
		"description of an event",
	)
	EventsCmd.Flags().StringVarP(
		&desc,
		"desc",
		"d",
		"",
		"description of an event",
	)
}
