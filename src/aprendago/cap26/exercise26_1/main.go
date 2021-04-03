/*
Aulas sobre tratamento de documentação de código go:
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=171
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=172
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=173
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=174
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=175
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=176
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=177
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=178
*/

//Package exercise26_1 contains functions that export metrics in the "key=value" format supported by monitoring tools like Prometheus, Zabbix, etc.
package exercise26_1

import "time"

// Func RecordMetrics generate metrics about processes.
func RecordMetrics() {
	go func() {
		for {
			time.Sleep(2 * time.Second)
		}
	}()
}

// Install godoc in Go 1.16.2 at Ubuntu 18.04
// go get -v  golang.org/x/tools/cmd/godoc

// Create doc of my modules
// godoc -http=:8080 -goroot ~/Documentos/mygit/go_mysamples
