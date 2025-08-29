package main

import "fmt"

func main() {
	ch := make(chan int, 4)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
		close(ch)
	}()

	for num := range ch {
		fmt.Println(num)
	}
}
