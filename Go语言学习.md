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

