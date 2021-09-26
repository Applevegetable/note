package main

import (
	"strconv"
	"strings"
)

//计算每个字母对应的数值
func key2value(key rune) int {
	var res int
	switch key{
		case 'G':
				res=6
		case 'O':
				res=14
		case 'L':
				res=11
		case 'A':
				res=0
		case 'N':
				res=13
	}		
	return res
}


func main() {
	msg:="csoiteuiwuiznsrocnkfd"
	var keyword  string = "GOLANG"
	msg=strings.ToUpper(msg)//转变为大写
	keyword =strings.ToUpper(keyword)//转换为大写

	for i:=0;i<len(msg);i++ {
	

	}
	


}