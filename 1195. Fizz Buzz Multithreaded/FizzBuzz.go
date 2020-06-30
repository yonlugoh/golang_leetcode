package main

// https://play.golang.org/p/22RA8BcVIYs
import (
	"fmt"
	"sync"
)

var (
	mutex sync.Mutex
	curr  int
)

func fizz(n int, wg *sync.WaitGroup) {
	for curr <= n {
		if curr%3 == 0 && curr%5 != 0 {
			mutex.Lock()
			fmt.Print("fizz,")
			curr++
			mutex.Unlock()
		}
	}
	wg.Done()
}

func buzz(n int, wg *sync.WaitGroup) {
	for curr <= n {
		if curr%3 != 0 && curr%5 == 0 {
			mutex.Lock()
			fmt.Print("buzz,")
			curr++
			mutex.Unlock()
		}
	}
	wg.Done()

}
func fizzbuzz(n int, wg *sync.WaitGroup) {
	for curr <= n {
		if curr%3 == 0 && curr%5 == 0 {
			mutex.Lock()
			fmt.Print("fizzbuzz,")
			curr++
			mutex.Unlock()
		}
	}
	wg.Done()

}

func number(n int, wg *sync.WaitGroup) {
	for curr <= n {
		if curr%3 != 0 && curr%5 != 0 {
			mutex.Lock()
			fmt.Print(curr, ",")
			curr++
			mutex.Unlock()
		}
	}
	wg.Done()

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
