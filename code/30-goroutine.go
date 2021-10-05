package main

import (
	"fmt"
	//"math/cmplx"
	"time"
)

func main() {
	c:=make(chan int)
	timeout := time.After(2 * time.Second)
	for i:=0;i<5;i++{
		select {
		case gopherID:=<-c:
			fmt.Println("gopher",gopherID,"has finished sleeping")
		case <-timeout:
			fmt.Println("my patience ran out")
		}
	}
}
func sleepyGopher(i int) {
	//time.Sleep(3*time.Second)

	fmt.Println("hello", i)

}
