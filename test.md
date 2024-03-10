<!-- do not add this code - testing -->

	app := tview.NewApplication()
	textView := tview.NewTextView().SetDynamicColors(true).SetRegions(true).SetWrap(true).SetText("GoVim: Mode-Based Editor\n ESC to exit\n")

	textArea := tview.NewTextArea().SetPlaceholder("Enter text here...")
	textArea.SetTitle("VimGo").SetBorder(false)
	helpInfo := tview.NewTextView().SetText("Press F1 for help, Ctrl-C to exit")
	position := tview.NewTextView().SetDynamicColors(true).SetTextAlign(tview.AlignRight)
	pages := tview.NewPages()

	updateInfos := func() {
		fromRow, fromColumn, toRow, toColumn := textArea.GetCursor()
		if fromRow == toRow && fromColumn == toColumn {
			position.SetText(fmt.Sprintf("Row: [yellow]%d[white], Column: [yellow]%d ", fromRow, fromColumn))
		} else {
			position.SetText(fmt.Sprintf("[blue]From[white] Row: [yellow]%d[white], Column: [blue]%d[white] - [red]To[white] Row: [yellow]%d[white], To Column: [yellow]%d ", fromRow, fromColumn, toRow, toColumn))
		}
	}

	textArea.SetMovedFunc(updateInfos)
	updateInfos()

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
