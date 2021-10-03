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

