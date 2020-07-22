package main

import (
	"dtn-gui/helpers"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	window, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)

	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	window.SetTitle("Data Transfer node")
	window.SetIconName("gtk-dialog-info")
	window.SetPosition(gtk.WIN_POS_CENTER)

	window.Connect("destroy", func() {
		println("got destroy!")
		gtk.MainQuit()
	})

	/* Global container vbox for all that's next */
	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	// box.PackStart(appLabel(), false, false, 0)

	/* Menu Bar */
	box.PackStart(helpers.BuildMenuBar(), false, false, 0)

	/* Frame for consumer*/
	frameC, _ := gtk.FrameNew("Consumer")
	box.PackStart(frameC, false, false, 0)

	frameC.Add(helpers.ConsumerFrame())

	/* Construct and show window */
	window.Add(box)
	window.SetDefaultSize(600, 400)
	window.ShowAll()
	gtk.Main()
}
