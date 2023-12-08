package main

import "fmt"

func main() {
	c1, c2 := make(chan int), make(chan string)
	go func() {
		for {
			select {
			case v := <-c1:
				fmt.Println("receive from c1:", v)
			case v := <-c2:
				fmt.Println("receive from c2:", v)
			}
		}
	}()
	c1 <- 1
	c2 <- "hello"
	c1 <- 2
	c2 <- "world"
}
