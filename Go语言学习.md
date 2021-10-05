# Go语言学习

## P2

1. Go编译器有哪些优点？
   - **1.自动垃圾回收降低了开发难度**
   - **2.更丰富的内置类型**，开发者**不用再费事添加依赖的包**，既减少了输入工作量，又可以让代码更简洁。
   - **3.支持函数多返回值**
   - **4.漂亮的错误处理**，Go语言引入了`defer`关键字用于标准的错误处理流程，并提供了内置函数panic、recover完成异常的抛出与捕获。与C++和Java等语言中的异常捕获机制相比，Go语言的错误处理机制可以大量减少代码量，开发者无需再仅为了程序安全而添加大量一层套一层的try-catch语句。
   - **5.匿名函数和闭包**,在Go语言中，所有的函数也是值类型，可以作为参数传递
   - **6.简洁的类型和“非侵入式”接口**
   - **7.并发编程更轻盈更安全**，goroutine是一种比线程更加轻盈、更省资源的协程。通过使用goroutine而不是裸用操作系统的并发机制，以及使用消息传递来共享内存而不是使用共享内存来通信，并发编程变得更加轻盈和安全。
   - 是一门编译型语言，在运行程序前，go首先使用编译器将你的代码转变为机器能读懂的1和0
   - 会把你所有的代码编译成一个可执行文件，在编译的过程中，编译器能捕获一些错误
   - 开源的软件
2. Go的程序从哪里运行？
   - package main包的main()函数
3. fmt这个package有哪些功能？
   - fmt包实现了类似C语言printf和scanf的格式化I/O。主要分为向外输出内容和获取输入内容两大部分
4. 左花括号放在哪里不会引起语法错误？
   - ​	func 同一行

## P3-计算器

Printf的第一个参数必须是字符串

- 这个字符串里面包含了%v这样的格式化动词，它的值由第二个参数的值所代替
- 如果指定了多个格式化动词，那么他们的值由后边的参数按顺序依次替代

使用Printf对齐文本

​	在格式化动词里指定宽度，就可以对齐文本

- **例如 %4v，就是向左填充到足够4个宽度**
  - **正数，向左填充空格，文本右对齐**
  - **负数，向右填充空格，文本左对齐**

常量和变量

- 常量const

- 变量var

- 多变量声明：

  ```
  var a,b=100,200
  var(
  	a=100
  	b=200
  )
  ```

- 赋值运算符

  - ```
    var weight=149.0
    weight + =200
    ```

- 自增运算符
  - ```go
    ++age 不可以
    age++ 可以
    ```

    

## P4  循环与分支

来自strings包的Contains函数可以判断某个字符串是否包含另外一个字符串



```go
	var command = "go inside"

	switch command {
	case "go test":
		fmt.Println("no")
	case "go inside", "go outside":
		fmt.Println("yes")
	default:
		fmt.Println("what")

	}
```

swith语句可以对数字进行匹配，默认有break关键字，不会执行下面的

还有一个fallthrough关键字，它用来执行下一个case的body部分

for不加关键字就是无线循环



作业：实现一个猜数游戏，首先定义一个1-100的整数，然后让计算机生成一个1-100的随机数，并显示计算机猜测的结果是太大了还是太小了，没猜对就继续猜，直到猜对为止。

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const res =79
	
	for {
		var random_num = rand.Intn(100)+1
		if random_num>res {
			fmt.Printf("%v is too big\n",random_num)
		}else if random_num<res {
			fmt.Printf("%v is too small\n",random_num)
		}else{
			fmt.Print("you are right,result is ",res)
			break
		}

	}
}
```

## P5 变量和作用域

**短声明**,有些时候没有办法使用var声明变量

```
p:=10
```

![image-20210925222958056](C:\Users\nature\AppData\Roaming\Typora\typora-user-images\image-20210925222958056.png)

![image-20210925223142445](C:\Users\nature\AppData\Roaming\Typora\typora-user-images\image-20210925223142445.png)

era变量的作用域是package main，对包内的其他函数都可见

短声明不可以用来声明package作用域的变量，就是说不可以`:=`这种方式只适用于函数内部，而不能成为全局变量

作业：

> •修改这个程序，让其能处理闰年的情况
>
> •生成随机年份，而不是写死 2018
>
> •二月份：闰年为 29 天，非闰年为 28 天
>
> •使用 for 循环生成和展示 10 个日期

![image-20210925223601880](C:\Users\nature\AppData\Roaming\Typora\typora-user-images\image-20210925223601880.png)

## P6 第一部分作业

![image-20210925230835340](C:\Users\nature\AppData\Roaming\Typora\typora-user-images\image-20210925230835340.png)

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const distance = 62100000 //距离火星距离

	fmt.Println("Spaceline            Days    Trip type     Price")
	fmt.Println("================================================")
	for num := 10; num > 0; num-- {
		var speed = rand.Intn(15) + 16 //速度

		var price = rand.Intn(15) + 36             //单程票价
		var days = distance / speed / 60 / 60 / 24 //日期
		var trip_type = rand.Intn(2)               //单程还是双程
		var spaceline = rand.Intn(3) + 1           //航线
		var line = ""
		var trip_name = ""
		var Price = 0
		if trip_type == 0 {
			trip_name = "One-way"
			Price = price
		} else {
			trip_name = "Round-way"
			Price = 2 * price
		}
		switch spaceline {
		case 1:
			line = "Space Adventures"
			fmt.Printf("%-19v  %-6v  %-12v  $  %v\n", line, days, trip_name, Price)
		case 2:
			line = "SpaceX"
			fmt.Printf("%-19v  %-6v  %-12v  $  %v\n", line, days, trip_name, Price)
		case 3:
			line = "Virgin Galactic"
			fmt.Printf("%-19v  %-6v  %-12v  $  %v\n", line, days, trip_name, Price)
		}

	}
}
```

## P7 实数

两种浮点数类型

默认是float64

- 64位的浮点类型
- 占用8字节内存
- 类似于double

float32

- 占用4字节内存
- 精度比float64低
- 有时叫做单精度

```go
var Pi  float64 = 63.92
```

![image-20210926100244758](C:\Users\hanyanbo\AppData\Roaming\Typora\typora-user-images\image-20210926100244758.png)

显示小数位数

```go
fmt.Printf("%.4f",third)
//小数位有4位
fmt.Printf("%4.2f",third)
//小数点2位，整体4位，包括小数点
```

![image-20210926101207111](C:\Users\hanyanbo\AppData\Roaming\Typora\typora-user-images\image-20210926101207111.png)

使用0代替空格进行填充

![image-20210926101242454](C:\Users\hanyanbo\AppData\Roaming\Typora\typora-user-images\image-20210926101242454.png)

浮点数不适合用于金融类计算，

为了尽量最小化舍入错误，建议先做乘法，再做除法

作业：

```go

/*
随机地将五分镍币（0.05美元）、一角硬币（0.10美元）和 25 美分硬币（0.25美元）放入一个空的储蓄罐，直到里面至少有20美元。
每次存款后显示存钱罐的余额
并以适当的宽度和精度格式化*/
package main

import (
	"fmt"
	"math/rand"
	
)

func main(){

	
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
```

## P8 整数

整数类型，包括有符号和无符号的

int 和uint是与操作系统有关系的，其他的都是无关的

**•而 int 和 uint 是针对目标设备优化的类型：**

•在树莓派 2、比较老的移动设备上，int 和 uint 都是 32 位的。

•在比较新的计算机上，int 和 uint 都是 64 位的。

**打印数据类型  %T**

•Go 语言里，在数前面加上 0x 前缀，就可以用十六进制的形式来表示数值。

•打印十六进制的数，使用 %x 格式化动词

![image-20210926105646341](C:\Users\hanyanbo\AppData\Roaming\Typora\typora-user-images\image-20210926105646341.png)

作业：

```go
/*
随机地将五分镍币（0.05美元）、一角硬币（0.10美元）和 25 美分硬币（0.25美元）放入一个空的储蓄罐，直到里面至少有20美元。
每次存款后显示存钱罐的余额
使用整数来追踪美分而不是美元
并以适当的宽度和精度格式化*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	const money = 2000

	//var coin float64
	var res int
	for res < money {
		num := rand.Intn(3)
		switch num {
		case 0:
			res += 5
			fmt.Printf("$%4.2f\n", float64(res)/100)

		case 1:
			res += 10
			fmt.Printf("$%4.2f\n", float64(res)/100)
		case 2:
			res += 25
			fmt.Printf("$%4.2f\n", float64(res)/100)

		}
	}

}
```

## P9 比较大的数

如果没有为指数形式的数值指定类型的话，那么Go会将它视为float64类型，

![image-20210926111200420](C:\Users\hanyanbo\AppData\Roaming\Typora\typora-user-images\image-20210926111200420.png)



在Go语言中，常量是可以无类型

•尽管 Go 编译器使用 big 包来处理无类型的数值常量，但是常量和 big.Int 的值是不能互换的。

使用big.Int来显示整数

```go
distance:= new (big.Int)
distance.SetString("24000000000000000000",10)
```

•尽管 Go 编译器使用 big 包来处理无类型的数值常量，但是常量和 big.Int 的值是不能互换的。

## P10 多语言文本

![image-20210926134825016](C:\Users\hanyanbo\AppData\Roaming\Typora\typora-user-images\image-20210926134825016.png)

![image-20210926135134140](C:\Users\hanyanbo\AppData\Roaming\Typora\typora-user-images\image-20210926135134140.png)

**自定义类型别名**

```go
type rune = int32
```

![image-20210926135448729](C:\Users\hanyanbo\AppData\Roaming\Typora\typora-user-images\image-20210926135448729.png)

打印字符本身，使用%c

打印字符对应的code points,使用%v,会自动推断其为rune类型



**string字符串本身是不可以改变的**

```go
var str string ="dsahdgas"
str[2]="x"//错误
```

• len 返回 message 所占的 byte 数。

•Go 有很多内置函数，它们不需要 import

如果是英语，返回的就是正确的长度，其他语言返回的字节数与长度不符合，因为Utf-8是可变长度的编码，汉字等会占据更大的空间





## P11 类型转换

![image-20210926145908032](C:\Users\hanyanbo\AppData\Roaming\Typora\typora-user-images\image-20210926145908032.png)

![image-20210926145921309](C:\Users\hanyanbo\AppData\Roaming\Typora\typora-user-images\image-20210926145921309.png)

![image-20210926154550426](C:\Users\hanyanbo\AppData\Roaming\Typora\typora-user-images\image-20210926154550426.png)

**fmt.Sprintf()**不会在控制台进行输出，转换到一个变量中进行输出

![image-20210926160232487](C:\Users\hanyanbo\AppData\Roaming\Typora\typora-user-images\image-20210926160232487.png)

像printf()一样，可以传入一个变量，%v等

**某些语言里，经常把 1 和 0 当作 true 和 false，但是在 Go 里面不行。**

```go

/*
写一个程序，把字符串转化为布尔类型：
“true”, “yes”, “1” 是 true
“false”, “no”, “0” 是 false
针对其它值，显示错误信息
*/
package main

import (
	"fmt"
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
```

## P13 函数

小写字母开头的函数不能被其他包调用，必须是**大写字母**开头的包

## P14 方法

**声明新类型 type celsius float64**

极大地提高代码的可读性和可靠性



Go语言中提供了方法，但是没有提供类和对象

每个方法可以有多个参数，但是只能有一个接收者

![image-20210926235728460](C:\Users\nature\AppData\Roaming\Typora\typora-user-images\image-20210926235728460.png)

```go
func (c celsius) kelvin() kelvin {

  return kelvin(c + 273.15)

}

//理解：receiver接收者是调用这个方法的类型，方法
```

## P15 一等函数

闭包就是由于匿名函数封闭并包围作用域中的变量而得名的

匿名函数就是函数名的函数

函数名可以直接传递进去







## P17数组

数组是一种固定长度且有序的元素集合

![image-20211001120552202](C:\Users\nature\AppData\Roaming\Typora\typora-user-images\image-20211001120552202.png)

•如果 Go 编译器在编译时未能发现越界错误，那么程序在运行时会出现 panic（例子）

•Panic 会导致程序崩溃

![image-20211001210508851](C:\Users\nature\AppData\Roaming\Typora\typora-user-images\image-20211001210508851.png)

range首先编译数组下标，之后遍历数组内的元素

## P18切片

切片不会导致数组被修改，只是创建了指向数组的一个窗口

slice是左闭右开的[0:4] 0,1,2,3

...表示可以自动推断出长度

直接声明切片

```go
dwarfs:=[]string ={"ceres","pluto","haumea","Makemake","Eris"}
```

切片类型的写法

```go
[]string
```

切片虽然有长度，但是这个长度与数组的长度不一样，它不是类型的一部分。可以将任意长度的切片传递给函数

**所有的range遍历第一个值都是下标，第二个值才是元素**

## P19更大的切片

append函数

```
append(切片，添加的元素)
```

容量和长度

如果切片底层的数组比切片大，那么就说该切片还有容量可以增长

#### 容量的理解

```go
slice1 := []string{"1", "2", "3", "4", "5"}
fmt.Println(len(slice1),cap(slice1))
slice2:=append(slice1,"6","8","8")
fmt.Println(len(slice2),cap(slice2))
fmt.Println(slice2)
slice3:=append(slice2,"9","10","12")
slice3[6]="7"
fmt.Println(slice2)
fmt.Println(slice3)
fmt.Println(len(slice3),cap(slice3))
	
	
5 5
8 10//新加入元素的个数超过底层的容量时，会新创建一个切片，容量是之前的两倍
[1 2 3 4 5 6 8 8]
[1 2 3 4 5 6 8 8]
[1 2 3 4 5 6 7 8 9 10 12]
11 20	
	
```

- **当新加入的元素是小于底层的容量的两倍时，会将容量改为两倍，即扩大一倍**
- 如果超过两倍，那么新的容量就是两次元素个数的加和
- 如果新加入元素之后容量还有剩余，那么就会共享底层的数组，进行修改，
- 如果没有剩余，就会新创建底层数组，

#### 三索引切分操作

```go
terrestrial：=plants[0:4:4]//从0-4，容量也是4
会创建新的切片

```

#### 使用make函数对切片实现预分配

```go
dwafs:=make([]string,0,10)//代表长度是0，容量是10
dwafs:=make([]string,10)//代表长度是10，容量是10,长度和容量一致，用一个参数来代替
//make创建出来的元素都是对应的零值
```

#### 声明可变参数函数

为了声明像Printf和append这样能够接受可变数量的实参的可变参数函数，我们需要在该函数的最后一个形参前面加上省略号...

通过省略号可以展开切片中的多个元素

#### GO语言中...的用法

1. 让Go编译器计算复合字面量中数组包含的元素个数
2. 创建可变参数函数的最后一个形参，使它可以将0个或者多个实参捕获为切片
3. 将切片中的元素展开为传递给函数的多个参数

## P20 MAP

![image-20211003141744516](C:\Users\nature\AppData\Roaming\Typora\typora-user-images\image-20211003141744516.png)

#### 逗号与ok语法

```go
	tempture := map[string]int{
		"earth": 15,
		"mars":  16,
	}

	if moon, ok := tempture["moon"]; ok {
		fmt.Print(moon)
		
	}else{
		fmt.Print()
	}
```

**用来判断是否存在键是否存在，如果存在ok被设置为true,注意有分号**

#### 映射不会被复制

数组在被赋值给新变量或者传递至函数或方法的时候都会创建对应的副本，但是映射不会，映射是直接对应的指针，指向的是相同的底层数据

#### 使用make预先分配空间

```go
temperature:=make(map[float64]int,8)//8代表有8个键值对
//使用make函数创建的新映射的初始长度总为0
```

**Go语言{}内必须使用，结尾**

## P22 结构

**声明结构体**

```go
type location struct{
		a int 
		b string
}
```

可以复用

**初始化**

```go
var s location //进行初始化
opp:=location{a:10,b:"sg"}//字面值初始化
opp1:=location{10,"sg"}//按照字段声明的顺序初始化
```

**打印struct**

![image-20211003222000026](C:\Users\nature\AppData\Roaming\Typora\typora-user-images\image-20211003222000026.png)

```go
%v 打印出的是{10,"sg"}
%+v:打印出的是{a:10,b:"sg"}
```

**结构体切片**

```go
	type location struct{
		name string
		lat  float64
		long float64
	}
	locations:=[]location{
		{name: "Baa",lat: -312,long: 29.2},
		{name: "Col",lat: 213.2,long: 212},
		{name: "Scs",lat: 212.3,long: 231},
	}
```

## P23-Go语言没有Class

**Go没有class,没有对象，也没有继承**

**但是提供了struct和方法，通过两者的组合就可以实现面向对象**

 

#### **构造函数的起名方式**

```go
newType//不可以被其他包引用
NewType//可以被外部引用
```

## P24组合与继承

Go只有组合没有继承

#### 结构嵌入

```
type report struct{
	sol int 
	tempture
	loccation//不给定字段名，只给定类型
}
```

可以将任意类型嵌入结构

## P25 接口

接口关注于类型可以做什么，而不是存储了什么

接口通过列举类型必须满足的一组方法来进行声明

在Go语言中，不需要显示声明接口

```go
var t interface{
	talk() string
}
type talker interface{
	talk() string
}
```

接口类型的名称通常以-er作为后缀

Go语言的接口是随时可以改变的，隐式实现

## P27指针

&	取地址，但是无法获得字符串、数值、布尔 字面值的地址

•&42，&“hello”这些都会导致编译器报错

\* 操作符与 & 的作用相反，它用来解引用，提供内存地址指向的值。

Go语言有指针，也强调安全性，不会出现野指针和迷途指针

 C 语言中的内存地址可以通过例如 address++ 这样的指针运算进行操作，但是在 Go 里面不允许这种不安全操作

**•将 * 放在类型前面表示声明指针类型**

**•将 * 放在变量前面表示解引用操作**

两个指针变量持有相同的内存地址，那么它们就是相等的



Go为结构体和数组实现了指针的自动解引用，但是对切片和映射map没有类似操作

Go语言的地址操作符&不仅可以获取结构的内存地址，还可以获取结构中指定字段的内存地址

映射和切片都是隐式指针

## P29-nil

尝试解引用一个nil指针会导致程序崩溃

## P30错误处理

error是一种类型

defer，延迟，Go可以确保所有defered的动作可以在函数返回前执行

在函数返回前肯定会触发defer

defer并不是专门做错误处理的

defer可以消除必须时刻惦记执行资源释放的负担

errors.New()函数生成一个错误，errors.New("this is a new error~")

自定义错误类型命名以Error结尾

panic不常使用

使用panic会执行所有defer动作

如果defer中调用了recover，那么panic就会停止，程序就会继续运行。

## P31goroutine

独立的任务叫做goroutine

启动goroutine使用关键字go

通道可以在多个goroutine之间安全的传值

创建通道,并指定其传输数据的类型

```
c:=make(chan int)
```

使用左箭头操作符<-向通道发送值或者从通道接受值

​	•	向通道发送值：c <- 99

​	•	从通道接收值：r := <- c

•发送操作会等待直到另一个 goroutine 尝试对该通道进行接收操作为止。

​	•执行发送操作的 goroutine 在等待期间将无法执行其它操作

​	•未在等待通道操作的 goroutine仍然可以继续自由的运行

•执行接收操作的 goroutine 将等待直到另一个 goroutine 尝试向该通道进行发送操作为止。

#### select处理多个通道

等待不同类型的值

time.After返回一个通道 ，该通道在指定时间后会接收到一个值 

```
select {
	case c:=<-channel:
	case b<-c
}
```

即使已经停止等待goroutine，但是只要main函数还没有返回，仍在运行的goroutine将会 继续占用内存

select在不包含任何goroutine的情况下会永远等下去

如果不适用make初始化通道，那么通道变量的值就是nil

•对 nil 通道进行发送或接收不会引起 panic，但会导致永久阻塞。

•对 nil 通道执行 close 函数，那么会引起 panic

•nil 通道的用处：

•对于包含 select 语句的循环，如果不希望每次循环都等待 select 所涉及的所有通道，那么可以先将某些通道设为 nil，等到发送值准备就绪之后，再将通道变成一个非 nil 值并执行发送操作。

#### 阻塞和死锁

•当 goroutine 在等待通道的发送或接收时，我们就说它被阻塞了。

•除了 goroutine 本身占用少量的内存外，被阻塞的 goroutine 并不消耗任何其它资源。

•goroutine 静静的停在那里，等待导致其阻塞的事情来解除阻塞。

•当一个或多个 goroutine 因为某些永远无法发生的事情被阻塞时，我们称这种情况为死锁。而出现死锁的程序通常会崩溃或挂起。  

•Go 允许在没有值可供发送的情况下通过 close 函数关闭通道

•例如 close(c)

•通道被关闭后无法写入任何值，如果尝试写入将引发 panic。

•尝试读取被关闭的通道会获得与通道类型对应的零值。

•注意：如果循环里读取一个已关闭的通道，并没检查通道是否关闭，那么该循环可能会一直运转下去，耗费大量 CPU 时间

•执行以下代码可得知通道是否被关闭：

•v, ok := <- c

## P35并发状态

互斥锁有两个方法Lock()和Unlock()

```
import "sync"
var mu sync.Mutex
func main(){
	mu.Lock()
	def mu.Unlock()
}
```

