package createac

import (
	"github.com/spf13/cobra"
	"tdahelper/cmd/createac/createsubac"
)

func init() {
	CreateCmd.AddCommand(createsubac.EventsCmd)
}

var CreateCmd = &cobra.Command{
	Use: "create",
}
