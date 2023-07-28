package bulksubac

import (
	"github.com/spf13/cobra"
	"tdahelper/config"
)

var file string

var EventsCmd = &cobra.Command{
	Use:   "event",
	Short: "bulk event -f file.csv",
	RunE: func(cmd *cobra.Command, args []string) error {
		return config.Mng.Bulk(file)
		//_ = render.RenderTable(td)
	},
}

func init() {
	EventsCmd.Flags().StringVarP(&file, "file", "f", "", "description of an event")
}
