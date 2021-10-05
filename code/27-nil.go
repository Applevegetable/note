package main

import (
	//"bufio"
	"fmt"
)

//物品
type item struct {
	name string
}

//角色
type character struct {
	name     string
	leftHand *item
}

//捡东西
func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}
	fmt.Printf("%v picks up a %v\n", c.name, i.name)
	c.leftHand = i
}

func (c *character) give(to *character) {
	if c == nil || to == nil {
		return
	}
	if c.leftHand == nil {
		fmt.Printf("%v has nothing to give\n", c.name)
		return 
	}
	if to.leftHand != nil {
		fmt.Printf("%v's hands are full\n", to.name)
		return 
	}
	to.leftHand=c.leftHand
	c.leftHand=nil
	fmt.Printf("%v give %v a %v\n", c.name, to.name, to.leftHand.name)
}
func (c character) String() string{
	if c.leftHand==nil{
		return fmt.Sprintf("%v is carrying nothing",c.name)
	}
	return fmt.Sprintf("%v is carrying a %v",c.name,c.leftHand.name)
}
func main() {
	arthu:=&character{name:"Arthur"}
	 
	shrubbery:= &item{name:"Shrubbery"}

	arthu.pickup(shrubbery)
	knight:=&character{name:"Knight"}
	arthu.give(knight)
	fmt.Println(arthu)
	fmt.Println(knight)
}
