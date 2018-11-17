package main // import "andlabsui"

import (
	"os"
	"time"

	"github.com/andlabs/ui"
	// _ "github.com/andlabs/ui/winmanifest"
)

func main() {
	err := ui.Main(func() {
		input := ui.NewEntry()
		button := ui.NewButton("Greet")
		greeting := ui.NewLabel("")

		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Enter your name:"), false)
		box.Append(input, false)
		box.Append(button, false)
		box.Append(greeting, false)

		window := ui.NewWindow("Hello", 200, 100, false)
		window.SetMargined(true)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			greeting.SetText("Hello, " + input.Text() + "!")
			time.Sleep(3000 * time.Millisecond)
			os.Exit(3)
		})

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
