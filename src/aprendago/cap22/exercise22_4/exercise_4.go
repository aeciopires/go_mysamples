/*
Aulas sobre canais:
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=146
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=147
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=148
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=149
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=150
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=151
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=152
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=153
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=154
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=155
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=156
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=157
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=158
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=159
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=160
*/

package main

import (
	"fmt"
)

func generateChannel(quitChan chan<- int, length int) <-chan int {
	channel := make(chan int)
	go func() {
		for i := 0; i < length; i++ {
			channel <- i
		}
		close(channel)
		quitChan <- 0
	}()

	return channel
}

func receiveChannel(channel <-chan int, quitChan chan int) {
	for {
		select {
		case value := <-channel:
			fmt.Println("Received of channel:", value)
		case <-quitChan:
			return
		}
	}
}

func main() {
	interactions := 100
	quitChannel := make(chan int)

	myChannel := generateChannel(quitChannel, interactions)
	receiveChannel(myChannel, quitChannel)

	fmt.Println("Exit...")
}
