
/*
写一个程序，把字符串转化为布尔类型：
“true”, “yes”, “1” 是 true
“false”, “no”, “0” 是 false
针对其它值，显示错误信息
*/
package main

import (
	"fmt"
	//"math"
)

func main(){
	var str string
	var res bool
	str="1"
	switch str {
		case "1", "true", "yes":
			res=true
			fmt.Print(res)
		case "0", "false", "no":
			res=false
			fmt.Print(res)
		default:
			fmt.Printf("this is wrong")
	}

	

}