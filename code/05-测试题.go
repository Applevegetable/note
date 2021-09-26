package main

import (
	"fmt"
	"math/rand"
)

func main4() {
	const distance = 62100000 //距离火星距离

	fmt.Println("Spaceline            Days    Trip type     Price")
	fmt.Println("================================================")
	for num := 10; num > 0; num-- {
		var speed = rand.Intn(15) + 16 //速度

		var price = rand.Intn(15) + 36             //单程票价
		var days = distance / speed / 60 / 60 / 24 //日期
		var trip_type = rand.Intn(2)               //单程还是双程
		var spaceline = rand.Intn(3) + 1           //航线
		var line = ""
		var trip_name = ""
		var Price = 0
		if trip_type == 0 {
			trip_name = "One-way"
			Price = price
		} else {
			trip_name = "Round-way"
			Price = 2 * price
		}
		switch spaceline {
		case 1:
			line = "Space Adventures"
			fmt.Printf("%-19v  %-6v  %-12v  $  %v\n", line, days, trip_name, Price)
		case 2:
			line = "SpaceX"
			fmt.Printf("%-19v  %-6v  %-12v  $  %v\n", line, days, trip_name, Price)
		case 3:
			line = "Virgin Galactic"
			fmt.Printf("%-19v  %-6v  %-12v  $  %v\n", line, days, trip_name, Price)
		}

	}
}
