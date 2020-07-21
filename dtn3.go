package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/gotk3/gotk3/gdk"
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
	box.PackStart(buildMenuBar(), false, false, 0)

	/* Frame for consumer*/
	frameC, _ := gtk.FrameNew("Consumer")
	box.PackStart(frameC, false, false, 0)

	frameC.Add(consumerFrame())

	/* Construct and show window */
	window.Add(box)
	window.SetDefaultSize(600, 400)
	window.ShowAll()
	gtk.Main()
}

func consumerFrame() *gtk.Box {

	boxC, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)

	/* a holder Hbox for consumer buttons */
	buttons, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	boxC.PackStart(buttons, false, false, 10)

	/* Buttons of frame 1 (Consumer) */
	button1, err := gtk.ButtonNewWithLabel("Initiatie")
	if err != nil {
		log.Fatal("Unable to create Button:", err)
	}
	button1.Connect("xxxxx ", func() {
		log.Println("test>>>>")
	})

	buttons.PackStart(button1, false, false, 10)
	// button2, _ := gtk.ButtonNewWithLabel("Start")
	// buttons.PackStart(button2, false, false, 10)
	// button3, _ := gtk.ButtonNewWithLabel("Status")
	// buttons.PackStart(button3, false, false, 10)
	// button4, _ := gtk.ButtonNewWithLabel("Abort")
	// buttons.PackStart(button4, false, false, 10)

	// // button1.Connect("initiatieConsumerbutton ", initiatieConsumerbutton)

	return boxC
}

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

func appLabel() *gtk.Label {
	label, _ := gtk.LabelNew("Data Transfer node")
	// label.ModifyFontEasy("DejaVu Serif 15")
	return label
}

func buildMenuBar() *gtk.MenuBar {
	menubar, _ := gtk.MenuBarNew()
	/* Menu bar item - File - */
	cascademenu, _ := gtk.MenuItemNewWithMnemonic("_File")
	menubar.Append(cascademenu)
	submenu, _ := gtk.MenuNew()
	cascademenu.SetSubmenu(submenu)

	/* Menu bar item - View - */
	cascademenu, _ = gtk.MenuItemNewWithMnemonic("_View")
	menubar.Append(cascademenu)
	submenu, _ = gtk.MenuNew()
	cascademenu.SetSubmenu(submenu)

	/* Menu bar item - Help - */
	cascademenu, _ = gtk.MenuItemNewWithMnemonic("_Help")
	menubar.Append(cascademenu)
	submenu, _ = gtk.MenuNew()
	cascademenu.SetSubmenu(submenu)

	helpItem, _ := gtk.MenuItemNewWithMnemonic("_About")
	submenu.Append(helpItem)

	helpItem.Connect("activate", func() {
		dialog, _ := gtk.AboutDialogNew()
		dialog.SetName("About DTN")
		dialog.SetProgramName(" ")
		pixbuf, _ := gdk.PixbufNewFromFile("/home/t3st/Desktop/2.png")
		dialog.SetLogo(pixbuf)
		dialog.SetAuthors(authors())
		dialog.SetLicense("The identification of any commercial product or trade name does not imply endorsement or recommendation by the National Institute of Standards and Technology, nor is it intended to imply that the materials or equipment identified are necessarily the best available for the purpose.")
		dialog.SetWrapLicense(true)
		dialog.Run()
		dialog.Destroy()
	})

	return menubar
}

func authors() []string {
	return []string{"Siham Khoussi <siham.khoussi@nist.gov>", "Abdella Battou <abdella.battou@nist.gov>"}
}

func server() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}

func client() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
