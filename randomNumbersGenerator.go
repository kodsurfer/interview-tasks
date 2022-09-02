package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for n := range randNumsGenerator(10) {
		fmt.Println(n)
	}
}

func randNumsGenerator(n int) <-chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- r.Intn(n)
		}
		close(out)
	}()

	return out
}
