package eventquery

import (
	"fmt"
	"github.com/spf13/cobra"
	"tdahelper/cmd/render"
	"tdahelper/config"
	"tdahelper/internal/model"
	"time"
)

var ByMonthQrCmd = &cobra.Command{
	Use: "bymonth",
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		now := time.Now()
		_, month, year := model.DateToDMY(now)
		return []string{fmt.Sprintf("%s-%s", year, month)}, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var td [][]string
		month, year, err := model.StrDateToMouthYear(args[0])
		if err != nil {
			return err
		}
		td = config.Mng.FindByMouth(month, year, cmd.Flags())
		return render.RenderTable(td)
	},
}
