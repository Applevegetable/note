package main

import "fmt"
//传入多个string类型的值，使用...的方式传入
func terraform(prefix string,worlds ...string)[]string{
	newWorlds:=make([]string,len(worlds))

	for i:=range worlds	{
		newWorlds[i]=prefix+" "+worlds[i]
	}
	return newWorlds
}
func main() {
	// slice1 := []string{"1", "2", "3", "4", "5"}
	// fmt.Println(len(slice1),cap(slice1))
	// slice2:=append(slice1,"6","8","8","6","8","8")
	// fmt.Println(len(slice2),cap(slice2))
	// fmt.Println(slice2)
	// slice3:=append(slice2,"9","10","12")
	// slice3[6]="7"
	// fmt.Println(slice2)
	// fmt.Println(slice3)
	// fmt.Println(len(slice3),cap(slice3))
	// twoWorlds:=terraform("New","Venus","Mars")
	// fmt.Println(twoWorlds)

	// plantes:=[]string{"Venus","Mars","Jupiter"}
	// newPlantes:=terraform("New",plantes...)
	// fmt.Println(newPlantes)


	s:=[]string{}
	lastCap:=cap(s)

	for i:=0;i<10000;i++{
		s=append(s,"An element")
		if cap(s)!=lastCap{
			fmt.Println(cap(s))
			lastCap=cap(s)
		}
	}

}