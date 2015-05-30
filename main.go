package main

import (
	"fmt"
	"math"

	"github.com/mgutz/logxi/v1"
)

var logger log.Logger

func main() {

	logger = log.New("sgmonitor")
	logger.Trace("I am about to call stuff")
	// logger.Debug("This is a debug log entry")
	// logger.Info("This is a info log entry")
	// logger.Warn("This is a warn log entry")
	// logger.Error("This is a error log entry")
	// logger.Fatal("This is a fatal log entry")
	// logger.log("This is a generic log entry")

	stuff(5)
	logger.Trace("I called stuff")
	fmt.Println(math.Pi)
}
