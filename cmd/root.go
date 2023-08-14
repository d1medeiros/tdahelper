package cmd

import (
	"github.com/spf13/cobra"
	"tdahelper/cmd/bulkac"
	"tdahelper/cmd/createac"
	"tdahelper/cmd/listac"
	"tdahelper/cmd/mainop"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string
	rootCmd     = &cobra.Command{
		Use:   "tda",
		Short: "A helper for adah",
		Long:  `A helper for adah with log description.`,
	}
)

func init() {
	rootCmd.AddCommand(listac.ListEventsCmd)
	rootCmd.AddCommand(mainop.OpCmd)
	rootCmd.AddCommand(createac.CreateCmd)
	rootCmd.AddCommand(bulkac.BulkCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
