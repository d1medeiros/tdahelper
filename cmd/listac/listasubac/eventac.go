package listasubac

import (
	"github.com/spf13/cobra"
	"tdahelper/cmd/listac/listasubac/eventquery"
	"tdahelper/cmd/render"
	"tdahelper/config"
	"tdahelper/internal/model"
)

var desc = ""
var category = ""
var sort = ""

var EventsCmd = &cobra.Command{
	Use:   "event",
	Short: "Print all the events based on filter",
	RunE: func(cmd *cobra.Command, args []string) error {
		var td [][]string
		td = config.Mng.FindAll(sort)
		return render.RenderTable(td)
	},
}

func init() {
	EventsCmd.AddCommand(eventquery.RangeQrCmd)
	EventsCmd.AddCommand(eventquery.ByMonthQrCmd)
	EventsCmd.PersistentFlags().StringVarP(&category, "category", "c", "", "")
	EventsCmd.PersistentFlags().StringVarP(&sort, "sort", "s", "", "")
	EventsCmd.RegisterFlagCompletionFunc(
		"sort",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return []string{model.EventsById, model.EventsByDate}, cobra.ShellCompDirectiveNoFileComp
		},
	)
	EventsCmd.RegisterFlagCompletionFunc(
		"category",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return config.Mng.FindAllCategory(), cobra.ShellCompDirectiveNoFileComp
		},
	)
}
