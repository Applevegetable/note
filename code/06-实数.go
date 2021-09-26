
/*
随机地将五分镍币（0.05美元）、一角硬币（0.10美元）和 25 美分硬币（0.25美元）放入一个空的储蓄罐，直到里面至少有20美元。
每次存款后显示存钱罐的余额
并以适当的宽度和精度格式化*/
package main

import (
	"fmt"
	"math/rand"
	
)

func main3(){

	
	const money=20
	//var coin float64
	var res  float64
	for res<money{
		num:=rand.Intn(3)
		switch num{
		case 0:
			res+=0.05
			fmt.Printf("余额为%4.2f\n",res)
		
	case 1:
		res+=0.10
		fmt.Printf("余额为%4.2f\n",res)
	case 2:
		res+=0.25
		fmt.Printf("余额为%4.2f\n",res)

		}
	}
	

}