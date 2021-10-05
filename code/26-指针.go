package main

import (
	"fmt"
	
)

type turtle struct {
	x, y int
}

func (t *turtle) up() {
	t.y--
}
func (t *turtle) down() {
	t.y++
}
func (t *turtle) left() {
	t.x--
}
func (t *turtle) right() {
	t.x++
}
func main() {
	t:=turtle{2,3}
	(&t).up()
	t.down()
	t.right()
	t.right()
	fmt.Println(t)
}
