package eventquery

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"tdahelper/cmd/render"
	"tdahelper/config"
	"tdahelper/internal/model"
	"time"
)

var RangeQrCmd = &cobra.Command{
	Use: "range",
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		now := time.Now()
		day, month, year := model.DateToDMY(now)
		return []string{fmt.Sprintf("%s-%s-%s", year, month, day)}, cobra.ShellCompDirectiveNoFileComp
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New(fmt.Sprintf("it is needed 2 args; args used %d", len(args)))
		}
		for i := 0; i < len(args); i++ {
			_, _, _, err := model.StrDateSplit(args[i])
			if err != nil {
				return err
			}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var td [][]string
		td = config.Mng.FindByRange(args[0], args[1], cmd.Flags())
		return render.RenderTable(td)
	},
}
