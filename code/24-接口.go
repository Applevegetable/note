package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	Name string     `json:"name"`
	Lat  cooridnate `json:"latitude"`
	Long cooridnate `json:"longitude"`
}

type cooridnate struct {
	d, m, s float64
	h       rune
}

func (c cooridnate) String() string {
	return fmt.Sprintf("%v°%v'%.1f\"%c", c.d, c.m, c.s, c.h)
}
func (l location) String() string {
	return fmt.Sprintf("%v,%v", l.Lat, l.Long)
}

func (c cooridnate) decimial() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1

	}
	return sign * (c.d + c.m/60 + c.s/3600)
}
func (c cooridnate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		DD  float64 `json:"decimal`
		DMS string  `json:"dms"`
		D   float64 `json:"degress"`
		M   float64 `json:"minutes"`
		S   float64 `json:"seconds"`
		H   string  `json:"hemisphere"`
	}{
		DD:  c.decimial(),
		DMS: c.String(),
		D:   c.d,
		M:   c.m,
		S:   c.s,
		H:   string(c.h),
	})
}
func main() {
	elysium := location{
		Name: "Elysium Planitia",
		Lat:  cooridnate{4, 30, 0.0, 'S'},
		Long: cooridnate{135, 54, 0, 'E'},
	}
	bytes, err := json.MarshalIndent(elysium, "", "")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(bytes))
}
