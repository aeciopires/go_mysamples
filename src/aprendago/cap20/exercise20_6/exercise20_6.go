/*
Aulas sobre concorrência, paralelismo, go routines:
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=125
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=126
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=127
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=128
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=129
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=130
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=139
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=140
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=141
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=142
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=143
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=144
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=145
*/

// - Run the program with commands:
//   - go run exercise20_6.go (output in terminal)
//   - go build exercise20_6.go (generate bin in actual directory)
//   - go install (generate bin in $USER/go/bin/)

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var myroutines int = 2
var wait sync.WaitGroup

func newroutines(routines int) {
	wait.Add(routines)
	for aux := 0; aux < routines; aux++ {
		go func(number int) {
			fmt.Println("Goroutine:", number+1)
			wait.Done()
		}(aux)
	}
}

func main() {
	fmt.Println("Operating System (GOOS):\t", runtime.GOOS)
	fmt.Println("Architecture (GOARCH):\t", runtime.GOARCH)
	fmt.Println("Number of CPU: ", runtime.NumCPU())
	fmt.Println("Number of Go routines in initial: ", runtime.NumGoroutine())
	newroutines(myroutines)
	fmt.Println("Number of Go routines in run time: ", runtime.NumGoroutine())
	wait.Wait()
}
