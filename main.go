package main

import (
	"fmt"
	"os"

	"github.com/c9s/goprocinfo/linux"
	"github.com/mgutz/logxi/v1"
)

var logger log.Logger

func main() {

	fmt.Println("SGMonitor starting up")
	logger = log.New("sgmonitor")
	logger.Trace("attempting to read /proc/cpuinfo")
	var realpath = "/proc/cpuinfo"
	var fakepath = "proc/cpuinfo"

	var procpath string

	//  Check to see if /proc/cpuinfo exists, if it does
	if _, err := os.Stat(realpath); os.IsNotExist(err) {
		procpath = fakepath
	} else {
		procpath = realpath
	}
	cpuinfo, err := linux.ReadCPUInfo(procpath)
	if err != nil {
		logger.Fatal("Unable to get cpuinfo", "ReadCPUInfo")
	}
	logger.Trace("read cpuinfo")

	logger.Info("We have this many cpus", "ReadCPUInfo", len(cpuinfo.Processors))
	logger.Info("We have this many Cores", "ReadCPUInfo", cpuinfo.NumCore())

	// fmt.Printf("%+v\n", cpuinfo)

}
