package main

import (
	"log"
	"os"
	"runtime"

	"github.com/gotk3/gotk3/gtk"
	"github.com/napsy/go-vte"
)

func init() {
	runtime.LockOSThread()

}
func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	win.SetTitle("protty")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	term, err := vte.TerminalNew()
	if err != nil {
		log.Fatal("Unable to create terminal:", err)
	}

	wd, _ := os.Getwd()

	term.SpawnAsyncSimple(wd, []string{"/bin/bash"}, []string{"TERM=xterm"})

	// Add the label to the window.
	win.Add(term)

	// Set the default window size.
	win.SetDefaultSize(800, 600)

	// Recursively show all widgets contained in this window.
	win.ShowAll()
	newSession()
	gtk.Main()
}
