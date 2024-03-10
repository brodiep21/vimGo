package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Mode int

const (
	NormalMode Mode = iota
	InsertMode
)

var currentMode Mode = NormalMode

func main() {
	// err := application.StartApplication()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	app := tview.NewApplication()

	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetText("Hello! Welcome to VimGO!\n")

	textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyESC:
			// Exit the application when the ESC key is pressed
			app.Stop()
		case tcell.KeyRune:
			// Display pressed keys
			textView.SetText(textView.GetText(true) + string(event.Rune()))
		}
		return event
	})

	textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch currentMode {
		case NormalMode:
			switch event.Rune() {
			case 'i': // 'i' to enter insert mode
				currentMode = InsertMode
				textView.SetText("Insert mode\n")
			}
			// Implement other Normal Mode keybindings here

		case InsertMode:
			switch event.Key() {
			case tcell.KeyESC:
				// Exit Insert Mode and go back to Normal Mode
				currentMode = NormalMode
				// Remove "-- INSERT --" indicator and update text view
				text := textView.GetText(true)
				textView.SetText(text[:len(text)-12]) // Adjust based on your indicator's length
			case tcell.KeyRune:
				// Insert the pressed key's character
				textView.SetText(textView.GetText(true) + string(event.Rune()))
			}
			// Implement other Insert Mode functionalities here
		}
		return event
	})

	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}
