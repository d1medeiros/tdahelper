package listasubac

import (
	"fmt"
	"github.com/spf13/cobra"
	"tdahelper/cmd/render"
	"tdahelper/config"
)

var CategoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Print the category from all events",
	Run: func(cmd *cobra.Command, args []string) {
		td := config.Mng.FindAllCategory()
		render.RenderLine(fmt.Sprintf("%v", td))
	},
}
