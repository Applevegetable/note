package main

import "fmt"

/*
func main() {
	c:= "L fdph L vdz L frqtxhuhg"
	for i:=0;i<len(c);i++{
		fmt.Printf("%c",c[i]-3)
	}

}
*/

func main(){
	str:="Hola Estación Espacial Internacional"
	for _,c := range str {
		
		if c>='a'&&c<='z'{
			c=c+13
			if(c>'z'){
				c=c-26
			}

		}
		if c>='A'&&c<='Z'{
			c=c+13
			if(c>'Z'){
				c=c-26
			}

		}
		fmt.Printf("%c",c)

	}
}