package main

import (
	"fmt"
	"time"
)

func foo(n int, oddChan chan bool, evenChan chan bool) {
	for i := 0; i < n; i++ {
		fmt.Print("foo")
		evenChan <- true
		<-oddChan
	}
}

func bar(n int, oddChan chan bool, evenChan chan bool) {
	for i := 0; i < n; i++ {
		<-evenChan
		fmt.Print("bar")
		oddChan <- true

	}
}

func fooBar(n int) {
	var oddChan chan bool = make(chan bool)
	var evenChan chan bool = make(chan bool)
	go foo(n, oddChan, evenChan)
	go bar(n, oddChan, evenChan)
	time.Sleep(2 * time.Second)
}
