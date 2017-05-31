package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("-----------------------")
	go say("hello")
	say("hello")
	fmt.Println("-----------------------")
}

func say(str string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(str)
	}
}
