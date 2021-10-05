package main

import (
	"fmt"
	//"go/printer"
	"math"
)

type location struct {
	lat, long float64
}
type world struct {
	radius float64
}
type cooridnate struct {
	d, m, s float64
	h       rune
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
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

var (
	mars  = world{radius: 3389.5}
	earth = world{radius: 6371}
)

func main() {
	spirit := newLocation(cooridnate{14, 34, 6.2, 'S'}, cooridnate{175, 28, 21.5, 'E'})
	opportunity := newLocation(cooridnate{1, 56, 46.3, 'S'}, cooridnate{354, 28, 24.2, 'E'})
	curiosity := newLocation(cooridnate{4, 35, 22.2, 'S'}, cooridnate{137, 26, 30.1, 'E'})
	insight := newLocation(cooridnate{4, 30, 0.0, 'S'}, cooridnate{135, 54, 0, 'E'})
	s:=[]float64{
		mars.distance(spirit,opportunity),
		mars.distance(spirit,curiosity),
		mars.distance(spirit,insight),
		mars.distance(opportunity,curiosity),
		mars.distance(insight,opportunity),
		mars.distance(curiosity,insight),
	}

	num_min:=s[0]
	num_max:=s[0]
	
	for _,s1 := range s{
		if s1>num_max{
			num_max=s1
		}
		if s1<num_min{
			num_min=s1
		}
	}
	fmt.Println("max_distance",num_max)
	fmt.Println("min_distance",num_min)
	london:=newLocation(cooridnate{51,30,0,'N'},cooridnate{0,8,0,'W'})
	paris:=newLocation(cooridnate{48,51,0,'N'},cooridnate{2,21,0,'E'})
	fmt.Println("london to paris",earth.distance(london,paris))
	
}
