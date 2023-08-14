package createsubac

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"tdahelper/config"
	"tdahelper/internal/model"
	"time"
)

var category = ""
var eventType = ""
var desc = ""
var value = ""
var dateStr = ""
var installment = 0

func init() {
	EventsCmd.Flags().IntVarP(&installment, "installment", "i", 0, "")
	EventsCmd.Flags().StringVarP(&category, "category", "c", "mercado", "")
	EventsCmd.Flags().StringVarP(&eventType, "type", "t", "unknown", "")
	EventsCmd.Flags().StringVarP(&dateStr, "date", "d", "", "-d 2023-01-01")
	EventsCmd.RegisterFlagCompletionFunc(
		"category",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return config.Mng.FindAllCategory(), cobra.ShellCompDirectiveNoFileComp
		},
	)
	EventsCmd.RegisterFlagCompletionFunc(
		"type",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return []string{"debit", "credit"}, cobra.ShellCompDirectiveNoFileComp
		},
	)
	EventsCmd.RegisterFlagCompletionFunc(
		"date",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			now := time.Now()
			today := model.DateToDMYString(now)
			return []string{today}, cobra.ShellCompDirectiveNoFileComp
		},
	)
}

var EventsCmd = &cobra.Command{
	Use:   "event",
	Short: "create event -c {category} -d {date} -t {type} [desc] [value]",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if eventType != "credit" && installment > 0 {
			return errors.New(fmt.Sprintf("%s can't have installment", eventType))
		}
		var date = dateStr
		_, err := model.StrToDate(date)
		if err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var desc = args[0]
		var value = args[1]
		date, _ := model.StrToDate(dateStr)
		return config.Mng.Create(
			desc,
			value,
			date,
			category,
			eventType,
			installment,
		)
	},
}
