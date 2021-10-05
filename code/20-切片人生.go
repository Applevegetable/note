package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

// 二维，切片类型，表示世界
type Universe [][]bool

//创建新的世界
func NewUniverse() Universe {
	u := make(Universe, height) //先是高度
	for i := range u {
		u[i] = make([]bool, width)//再是宽度
	}
	return u
}

//打印世界状态
func (u Universe) Show() {
	fmt.Print("\x0c",u.String())
}
func (u Universe)String() string{
	var b byte
	buf:=make([]byte,0,(width+1)*height)//多打印一列,容量大一些

	for y:=0;y<height;y++{
		for x:=0;x<width;x++{
			b='c'
			if u[y][x]{
				b='*'
			}
			buf=append(buf, b)
		}
			buf=append(buf, '\n')
		}
		return string(buf)
}

//随机激活世界中25%的细胞
func (u Universe) Seed() {
	for i:=0;i<(width*height/4);i++{
			y:=rand.Intn(height)
			x:=rand.Intn(width)
			u[y][x]=true			
	}
}
//判断是否存活，世界将进行回绕，比如16（15）就会变为1，-1就会变为14
func (u Universe) Alive(x int,y int) bool{
	x=(x+width)%width
	y=(y+height)%height
	return u[y][x]
}

//统计临近细胞数量
func (u Universe) Neighbors(x,y int) int{
	num:=0
	for v:=-1;v<=1;v++{
		for h:=-1;h<=1;h++{
			if (v!=0&&h!=0)&& u.Alive(x+h,y+v){
				num++
			}
		}
	}
	return num
}
//游戏逻辑
func (u Universe) Next(x,y int)bool{
	n:=u.Neighbors(x,y)
	status:=u[y][x]
	//活着的细胞
	if status {
		if n==2||n==3{//只有周围只有两到三个活着的才能存活
			status=true
		}else{
			status=false
		}
	}else{//死掉的细胞
		if n==3{//只有临近有3个细胞存活才可以生存
			status=true
		}else{
			status=false
		}
	}
	return status
}

//平行世界
func Step(a,b Universe){
	for y:=0;y<height;y++{
		for x:=0;x<width;x++{
			b[y][x]=a.Next(x,y)
		}
	}
}


func main() {
	a ,b:= NewUniverse(),NewUniverse()//创造新的世界
	a.Seed()
	for i:=0;i<30;i++{
		Step(a,b)
		a.Show()
		time.Sleep(time.Second)
		a,b=b,a//平行宇宙交互显示与操作

	}
	
}
