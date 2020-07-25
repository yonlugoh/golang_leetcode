package main

// https://play.golang.org/p/7nuF5TZEiWM

import (
	"fmt"
	"sync"
)

var (
	mutex sync.Mutex
	curr  int
)

func fizz(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		mutex.Lock()
		if curr > n {
			mutex.Unlock()
			return
		}
		if curr%3 == 0 && curr%5 != 0 {
			fmt.Print("fizz,")
			curr++
		}
		mutex.Unlock()
	}

}

func buzz(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		mutex.Lock()
		if curr > n {
			mutex.Unlock()
			return
		}
		if curr%3 != 0 && curr%5 == 0 {
			fmt.Print("buzz,")
			curr++
		}
		mutex.Unlock()
	}

}

func fizzbuzz(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		mutex.Lock()
		if curr > n {
			mutex.Unlock()
			return
		}
		if curr%3 == 0 && curr%5 == 0 {
			fmt.Print("fizzbuzz,")
			curr++
		}
		mutex.Unlock()
	}

}

func number(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		mutex.Lock()
		if curr > n {
			mutex.Unlock()
			return
		}
		if curr%3 != 0 && curr%5 != 0 {
			fmt.Print(curr, ",")
			curr++
		}
		mutex.Unlock()
	}

}
func FizzBuzzMultithreaded(n int) {
	curr = 1
	var wg sync.WaitGroup
	wg.Add(4)
	go fizz(n, &wg)
	go buzz(n, &wg)
	go fizzbuzz(n, &wg)
	go number(n, &wg)
	wg.Wait()
}
