package main

import "fmt"

type celsius float64

//c->k
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

//c->f
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32)
}

type kelvin float64

//k->c
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

//k->f
func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

type fahrenheit float64

// f->c
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

//f->k
func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

func main() {

	var k kelvin = 294.0
	c := k.celsius()
	fmt.Print(k, c)

}
