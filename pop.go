package pop

import (
	"log"
	"os"
)

var AvailableDialects = []string{"postgres", "mysql", "cockroach"}

// Debug mode, to toggle verbose log traces
var Debug = false

// Color mode, to toggle colored logs
var Color = true
var logger = log.New(os.Stdout, "[POP] ", log.LstdFlags)

// Log a formatted string to the logger
var Log = func(s string, args ...interface{}) {
	if Debug {
		Logger.Info(s, args)
	}
}
