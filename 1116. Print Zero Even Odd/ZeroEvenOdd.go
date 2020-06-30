package _116__Print_Zero_Even_Odd

import (
	"fmt"
	"sync"
)

func ZeroEvenOdd(count int) {
	var wg sync.WaitGroup
	printC := make(chan int, 1)
	triggerOdd := make(chan struct{}, 1)
	triggerEven := make(chan struct{}, 1)
	wg.Add(3)
	go zero(&wg, printC, triggerEven, triggerOdd, count)
	go odd(&wg, printC, triggerOdd, count)
	go even(&wg, printC, triggerEven, count)
	printC <- 2
	wg.Wait()
}

func zero(wg *sync.WaitGroup, printC chan int, triggerEven chan struct{}, triggerOdd chan struct{}, count int) {
	for i := 0; i < count; i++ {
		p := <-printC
		fmt.Print("0")
		if p == 2 {
			triggerOdd <- struct{}{}
		} else {
			triggerEven <- struct{}{}
		}
	}
	wg.Done()
}

func odd(wg *sync.WaitGroup, printC chan int, triggerOdd chan struct{}, count int) {
	for i := 1; i <= count; i += 2 {
		<-triggerOdd
		fmt.Print(i)
		printC <- 1
	}
	wg.Done()
}

func even(wg *sync.WaitGroup, printC chan int, triggerEven chan struct{}, count int) {
	for i := 2; i <= count; i += 2 {
		<-triggerEven
		fmt.Print(i)
		printC <- 2
	}
	wg.Done()
}
