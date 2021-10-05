package main

import (
	"fmt"
	
)

//声明切片类型
type lantes []string

func (pla lantes)terraform() {
	for i:=range pla{
		pla[i]="New "+pla[i]
	}
}

func main(){
	plantes:=[]string{"Mercury","Venus","Earth","Mars","Jupiter","Staurn","Uranus","Neptune"}
	/*
	//切片类型
	fmt.Printf("%T",plantes[3:4])
	//string类型
	fmt.Printf("%T",plantes[3])
	*/
	earth1:=plantes[3:4]
	earth2:=plantes[6:]
	lantes(earth1).terraform()
	lantes(earth2).terraform()
	fmt.Print(plantes)

}