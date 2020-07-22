package main

import (
	"errors"
	"log"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// var builder *gtk.Builder

func errorCheck(e error) {
	if e != nil {
		// panic for any errors.
		log.Panic(e)
	}
}

func on_button_clicked() {
	log.Println("Button clicked")
}

func isWindow(obj glib.IObject) (*gtk.Window, error) {
	// Make type assertion (as per gtk.go).
	if win, ok := obj.(*gtk.Window); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.Window")
}

func isButton(obj glib.IObject) (*gtk.Button, error) {
	// Make type assertion (as per gtk.go).
	if button, ok := obj.(*gtk.Button); ok {
		return button, nil
	}
	return nil, errors.New("not a *gtk.Button")
}

func main() {
	gtk.Init(nil)
	// // Get the GtkBuilder UI definition in the glade file.
	builder, err := gtk.BuilderNewFromFile("./example.glade")
	errorCheck(err)

	// Map the handlers to callback functions, and connect the signals to the Builder.
	signals := map[string]interface{}{
		"on_button_clicked": on_button_clicked,
	}
	builder.ConnectSignals(signals)

	// Get the object with the id of "window".
	win, err := builder.GetObject("window")
	errorCheck(err)

	// Verify that the object is a pointer to a gtk.ApplicationWindow.
	window, err := isWindow(win)
	errorCheck(err)

	window.Connect("destroy", func() {
		println("got destroy!")
		gtk.MainQuit()
	})

	// Get the object with the id of "fixed".
	_, err = builder.GetObject("fixed")
	errorCheck(err)

	// Get the object with the id of "fixed".
	but, err := builder.GetObject("button")
	errorCheck(err)

	button, err := isButton(but)
	errorCheck(err)

	cssProv, err := gtk.CssProviderNew()
	cssProv.LoadFromPath("./modified.css")
	errorCheck(err)

	context, err := button.GetStyleContext()
	context.AddProvider(cssProv, gtk.STYLE_PROVIDER_PRIORITY_USER)
	context.AddClass("button")

	// Show the Window and all of its components.
	window.SetDefaultSize(600, 400)
	window.ShowAll()
	gtk.Main()

}
