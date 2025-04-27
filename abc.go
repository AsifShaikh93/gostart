package main

import (
	"fmt"
)

func Goprac() {
	ch1 := make(chan string)
	go func() {
		ch1 <- "heyy i am channel data"
	}()
	fmt.Println(<-ch1)
}
