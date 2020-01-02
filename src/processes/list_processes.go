/*
References:
https://stackoverflow.com/questions/9030680/list-of-currently-running-process-in-go
https://github.com/shirou/gopsutil
http://tdoc.info/en/blog/2015/12/16/gopsutil.html
https://ispycode.com/GO/System-and-Process-Utilities/CPU-Usage
https://github.com/tools/godep/issues/377

Resolve dependences:

go get github.com/mitchellh/go-ps
go get github.com/shirou/gopsutil
*/

// Definition of my package
package main

// Importing third packages
import (
	"fmt"

	ps "github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/mem"
)

// Implementing my funtions
func main() {
	// Create and initialize variables for process
	processList, err := ps.Processes()

	// Handling exception
	if err != nil {
		fmt.Printf("ps.Processes() Failed, are you using windows?")
		return
	}

	// Index process list
	for aux := range processList {
		var process ps.Process
		process = processList[aux]
		fmt.Printf("PID: %d --> Name: %s --> ParentPID: %d\n", process.Pid(), process.Executable(), process.PPid())
	}

	// Create and initialize variables for memory
	vmemory, err := mem.VirtualMemory()
	fmt.Printf("Total: %v --> Free: %v --> UsedPercent: %f%%\n", vmemory.Total, vmemory.Free, vmemory.UsedPercent)

	// Show complete result
	fmt.Println(vmemory)
}
