package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, n := range []int{1, 2, 3} {
			a <- n
		}
		close(a)
	}()

	go func() {
		for _, n := range []int{4, 5, 6} {
			b <- n
		}
		close(b)
	}()

	go func() {
		for _, n := range []int{7, 8, 9} {
			c <- n
		}
		close(c)
	}()

	for n := range joinChannels(a, b, c) {
		fmt.Println(n)
	}
}

func joinChannels(chs ...chan int) <-chan int {
	res := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(len(chs))

		for _, ch := range chs {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for c := range ch {
					res <- c
				}
			}(ch, wg)
		}
		wg.Wait()
		close(res)
	}()

	return res
}
