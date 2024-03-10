package main

import "github.com/rivo/tview"

func main() {
	app := tview.NewApplication()
	textView := tview.NewTextView().SetDynamicColors(true).SetRegions(true).SetWrap(true).SetText("GoVim: Mode-Based Editor\n ESC to exit\n")

	textList := tview.NewList().
		AddItem("List Item 1", "Explanation", 'a', nil).
		AddItem("List Item 1", "Explanation", 'b', nil).
		AddItem("List Item 1", "Explanation", 'c', nil).
		AddItem("List Item 1", "Explanation", 'd', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
	if err := app.SetRoot(textList, true).SetFocus(textList).Run(); err != nil {
		panic(err)
	}

}
