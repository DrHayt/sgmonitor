package main

import (
	"fmt"
	"math"

	"github.com/mgutz/logxi/v1"
)

func main() {

	logger := log.New("mylog")

	logger.Trace("This is a trace log entry", "main", "foo")
	logger.Trace("This is a trace log entry", "debug", "foo")
	logger.Trace("This is a trace log entry", "app", "foo")
	logger.Trace("This is a trace log entry", "sub", "foo")
	logger.Trace("This is a trace log entry", "foo", "foo")
	// logger.Debug("This is a debug log entry")
	// logger.Info("This is a info log entry")
	// logger.Warn("This is a warn log entry")
	// logger.Error("This is a error log entry")
	// logger.Fatal("This is a fatal log entry")
	// logger.log("This is a generic log entry")

	fmt.Println(math.Pi)
}
