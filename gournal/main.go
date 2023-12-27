package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func usercmd(t chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		t <- scanner.Text()
	}

}

func doSmth(cmd chan string) {
	for i := 0; i < 10; i++ {
		select {
		case <-cmd:
			break
		}
		fmt.Print(i)
		time.Sleep(time.Second)
	}

}

func main() {
	tunnel := make(chan string)
	for {
		go usercmd(tunnel)
		doSmth(tunnel)
		// fmt.Println("Somethings happening...", <-tunnel)
	}
}
