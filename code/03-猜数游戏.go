package main

import (
	"fmt"
	"math/rand"
)

func main1() {
	const res =79
	
	for {
		var random_num = rand.Intn(100)+1
		if random_num>res {
			fmt.Printf("%v is too big\n",random_num)
		}else if random_num<res {
			fmt.Printf("%v is too small\n",random_num)
		}else{
			fmt.Print("you are right,result is ",res)
			break
		}

	}



}
