package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func newSession() {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	win.SetTitle("New Session")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	buttonOK, _ := gtk.ButtonNewWithLabel("Ok")
	buttonCancel, _ := gtk.ButtonNewWithLabel("Cancel")

	labelName, _ := gtk.LabelNew("Name")
	labelHost, _ := gtk.LabelNew("Hostname")
	labelUser, _ := gtk.LabelNew("Username")
	labelPasswd, _ := gtk.LabelNew("Password")

	entryName, _ := gtk.EntryNew()
	entryHost, _ := gtk.EntryNew()
	entryUser, _ := gtk.EntryNew()
	entryPasswd, _ := gtk.EntryNew()

	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 2)
	box.PackStart(labelName, false, false, 2)
	box.PackStart(entryName, false, false, 2)
	box.PackStart(labelHost, false, false, 2)
	box.PackStart(entryHost, false, false, 2)
	box.PackStart(labelUser, false, false, 2)
	box.PackStart(entryUser, false, false, 2)
	box.PackStart(labelPasswd, false, false, 2)
	box.PackStart(entryPasswd, false, false, 2)
	box.SetBorderWidth(4)

	boxButtons, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 2)
	boxButtons.PackEnd(buttonOK, false, false, 2)
	boxButtons.PackEnd(buttonCancel, false, false, 2)

	buttonOK.Connect("clicked", func() {
		gtk.MainQuit()
	})
	buttonCancel.Connect("clicked", func() {
		gtk.MainQuit()
	})

	box.PackStart(boxButtons, false, false, 2)

	win.SetDefaultSize(200, 200)

	win.Add(box)
	win.ShowAll()
}
