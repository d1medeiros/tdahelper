package createsubac

import (
	"github.com/spf13/cobra"
	"tdahelper/config"
	"tdahelper/internal/model"
)

var EventsCmd = &cobra.Command{
	Use:   "event",
	Short: "create event [desc] [value] [date] [category]",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		var date = args[2]
		_, err := model.StrToDate(date)
		if err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var desc = args[0]
		var value = args[1]
		var dateStr = args[2]
		date, _ := model.StrToDate(dateStr)
		var category = args[3]
		return config.Mng.Create(desc, value, date, category)
		//_ = render.RenderTable(td)
	},
}
