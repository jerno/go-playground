package astilectron

import (
	"fmt"
	"log"
	"os"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

func Run() {
	// Set logger
	logger := log.New(log.Writer(), log.Prefix(), log.Flags())

	// Create astilectron
	application, err := astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		AppName:           "Erno's GO launcher",
		BaseDirectoryPath: "base-dir",
	})
	if err != nil {
		logger.Fatal(fmt.Errorf("main: creating astilectron failed: %w", err))
	}
	defer application.Close()

	// Handle signals
	application.HandleSignals()

	// Start
	if err = application.Start(); err != nil {
		logger.Fatal(fmt.Errorf("main: starting astilectron failed: %w", err))
	}

	// New window
	var window *astilectron.Window
	if window, err = application.NewWindow("examples/astilectron/resources/index.html", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(800),
		Width:  astikit.IntPtr(1200),
	}); err != nil {
		logger.Fatal(fmt.Errorf("main: new window failed: %w", err))
	}

	// Create windows
	if err = window.Create(); err != nil {
		logger.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	}

	// This will listen to messages sent by Javascript
	window.OnMessage(func(eventMessage *astilectron.EventMessage) interface{} {
		// Unmarshal
		var message string
		eventMessage.Unmarshal(&message)

		if message == "hello GO, it's JS" {
			return "hello JS, it's GO"
		} else if message == "OPEN_DEV_TOOLS" {
			window.OpenDevTools()
		}
		return nil
	})

	// This will send a message and execute a callback
	// Callbacks are optional
	window.SendMessage("Are you there, JS?", func(eventMessage *astilectron.EventMessage) {
		// Unmarshal
		var message string
		eventMessage.Unmarshal(&message)

		// Process message
		log.Printf("received: %s\n", message)
	})

	// Blocking pattern
	application.Wait()
}
