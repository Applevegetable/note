package main

import "fmt"

/**/
type location struct {
	lat, long float64
}
type cooridnate struct {
	d, m, s float64
	h       rune
}

//时分秒转换为十进制
func (c cooridnate) decimial() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1

	}
	return sign * (c.d + c.m/60 + c.s/3600)
}
func newLocation(lat, long cooridnate) location {
	return location{lat.decimial(), long.decimial()}
}
func main() {
	spirit := newLocation(cooridnate{14, 34, 6.2, 'S'}, cooridnate{175, 28, 21.5, 'E'})
	opportunity := newLocation(cooridnate{1, 56, 46.3, 'S'}, cooridnate{354, 28, 24.2, 'E'})
	curiosity := newLocation(cooridnate{4, 35, 22.2, 'S'}, cooridnate{137, 26, 30.1, 'E'})
	insight := newLocation(cooridnate{4, 30, 0.0, 'S'}, cooridnate{135, 54, 0, 'E'})
	fmt.Println("Spirit", spirit)
	fmt.Println("Opportunity", opportunity)
	fmt.Println("Curiosity", curiosity)
	fmt.Println("Insight", insight)
}
