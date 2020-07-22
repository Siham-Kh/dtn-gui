package helpers

import (
	"github.com/gotk3/gotk3/gtk"
)

func fill(progress *gtk.ProgressBar, fraction float64) bool {

	/*Fill in the bar with the new fraction*/
	progress.SetFraction(fraction)

	/*Ensures that the fraction stays below 1.0*/
	if fraction < 1.0 {
		return true
	} else {
		return false
	}
}

func AppLabel() *gtk.Label {
	label, _ := gtk.LabelNew("Data Transfer node")
	// label.ModifyFontEasy("DejaVu Serif 15")
	return label
}

func authors() []string {
	return []string{"Siham Khoussi <siham.khoussi@nist.gov>", "Abdella Battou <abdella.battou@nist.gov>"}
}

func initiatieConsumerbutton() {

	println("Initiating the client... ")
	window, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Starting the transfer...")

	notebook, _ := gtk.NotebookNew()

	fixed, _ := gtk.FixedNew()
	label, _ := gtk.LabelNew("Fixed")
	notebook.AppendPage(fixed, label)

	button, _ := gtk.ButtonNewWithLabel("Pulse")
	fixed.Put(button, 30, 30)

	progress, _ := gtk.ProgressBarNew()
	progress.SetOrientation(gtk.ORIENTATION_HORIZONTAL)
	fixed.Put(progress, 100, 100)

	var fraction float64 = 0
	/*Get the current progress*/
	fraction = progress.GetFraction()

	/*Increase the bar by 10% each time this function is called*/
	step := 0.1

	button.Connect("clicked", func() {
		// client()
		fraction += step
		fill(progress, fraction)
	})

	window.Add(notebook)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetDefaultSize(400, 200)
	window.ShowAll()
}

func ConsumerFrame() *gtk.Box {

	boxC, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)

	/* a holder Hbox for consumer buttons */
	buttons, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	boxC.PackStart(buttons, false, false, 10)

	/* Buttons of frame 1 (Consumer) */
	button1, _ := gtk.ButtonNewWithLabel("Initiatie")
	buttons.PackStart(button1, false, false, 10)
	button2, _ := gtk.ButtonNewWithLabel("Start")
	buttons.PackStart(button2, false, false, 10)
	button3, _ := gtk.ButtonNewWithLabel("Status")
	buttons.PackStart(button3, false, false, 10)
	button4, _ := gtk.ButtonNewWithLabel("Abort")
	buttons.PackStart(button4, false, false, 10)

	button1.Connect("clicked", func() {
		initiatieConsumerbutton()
	})

	return boxC
}
