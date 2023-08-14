package render

import "github.com/pterm/pterm"

func RenderTable(td [][]string) error {
	return pterm.DefaultTable.
		WithHasHeader().
		WithHeaderRowSeparator(".").
		WithBoxed().
		WithRowSeparator(".").
		WithData(td).
		Render()
}

func RenderLine(txt string) {
	pterm.DefaultBasicText.Println(txt)
}
