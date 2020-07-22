package helpers

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func BuildMenuBar() *gtk.MenuBar {
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
