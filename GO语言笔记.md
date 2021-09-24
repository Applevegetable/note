# GO语言笔记

[TOC]

vscode 返回上一个浏览位置

Alt+左右方向键



`runtime` 调度器是个非常有用的东西，关于 `runtime` 包几个方法:

- **Gosched**：让当前线程让出 `cpu` 以让其它线程运行,它不会挂起当前线程，因此当前线程未来会继续执行
- **NumCPU**：返回当前系统的 `CPU` 核数量
- **GOMAXPROCS**：设置最大的可同时使用的 `CPU` 核数
- **Goexit**：退出当前 `goroutine`(但是`defer`语句会照常执行)
- **NumGoroutine**：返回正在执行和排队的任务总数
- **GOOS**：目标操作系统

## NumCPU



```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println("cpus:", runtime.NumCPU())
    fmt.Println("goroot:", runtime.GOROOT())
    fmt.Println("archive:", runtime.GOOS)
}
```

运行结果：

![img](https:////upload-images.jianshu.io/upload_images/1262158-2fd4d08e60117ee4.png?imageMogr2/auto-orient/strip|imageView2/2/w/414/format/webp)

## GOMAXPROCS

`Golang` 默认所有任务都运行在一个 `cpu` 核里，如果要在 `goroutine` 中使用多核，可以使用 `runtime.GOMAXPROCS` 函数修改，当参数小于 1 时使用默认值。



```go
package main

import (
    "fmt"
    "runtime"
)

func init() {
    runtime.GOMAXPROCS(1)
}

func main() {
    // 任务逻辑...

}
```

## Gosched

这个函数的作用是让当前 `goroutine` 让出 `CPU`，当一个 `goroutine` 发生阻塞，`Go` 会自动地把与该 `goroutine` 处于同一系统线程的其他 `goroutine` 转移到另一个系统线程上去，以使这些 `goroutine` 不阻塞



```go
package main

import (
    "fmt"
    "runtime"
)

func init() {
    runtime.GOMAXPROCS(1)  //使用单核
}

func main() {
    exit := make(chan int)
    go func() {
        defer close(exit)
        go func() {
            fmt.Println("b")
        }()
    }()

    for i := 0; i < 4; i++ {
        fmt.Println("a:", i)

        if i == 1 {
            runtime.Gosched()  //切换任务
        }
    }
    <-exit

}
```

结果：

![img](https:////upload-images.jianshu.io/upload_images/1262158-80b91ce8a00bea07.png?imageMogr2/auto-orient/strip|imageView2/2/w/423/format/webp)

使用多核测试：



```go
package main

import (
    "fmt"
    "runtime"
)

func init() {
    runtime.GOMAXPROCS(4)  //使用多核
}

func main() {
    exit := make(chan int)
    go func() {
        defer close(exit)
        go func() {
            fmt.Println("b")
        }()
    }()

    for i := 0; i < 4; i++ {
        fmt.Println("a:", i)

        if i == 1 {
            runtime.Gosched()  //切换任务
        }
    }
    <-exit

}
```

结果：

![img](https:////upload-images.jianshu.io/upload_images/1262158-aab5e17653dfac6e.png?imageMogr2/auto-orient/strip|imageView2/2/w/363/format/webp)

根据你机器来设定运行时的核数，但是运行结果不一定与上面相同，或者在 `main` 函数的最后加上 **select{}**  让程序阻塞，则结果如下：

![img](https:////upload-images.jianshu.io/upload_images/1262158-278fe3220e4e27f5.png?imageMogr2/auto-orient/strip|imageView2/2/w/553/format/webp)

多核比较适合那种 `CPU` 密集型程序，如果是 `IO` 密集型使用多核会增加 `CPU` 切换的成本。





切片的容量大于等于长度

先定义长度，后定义容量

```go
slice:=make([]string,3,5)//长度为3，容量为5
```

切片字面量

初始的长度和容量会基于初始化时提供的元素的个数确定

创建一个整形切片，长度和容量都是3

```
slice:=[]int{10，20，30}
```



如果在[]运算符里指定一个值，那么创建的就是数组而不是切片。只有不指定的时候，才会创建切片。

```
array := [3]int{10,20,30}//有3个元素的数组
array := []int{10,20,30}//创建长度和容量都是3的切片
```

创建nil切片

```go
var slice []int 
```

利用初始化，通过声明一个切片创建一个空切片

```
slice := make([]int ,0)//使用make创建空的整型切片
slice := []int{}//使用切片字面量创建空的整型切片
```

切片之所以称为切片，是因为创建一个新的切片就是把底层数组切出一部分

```go
slice :=[]int{10,20,30,40,50}//创建一个整型切片，长度和容量都是5
newSlice:=slice[1:3]//创建一个新切片，其长度为2，容量为4
```

切片只能访问到其长度内的元素，试图访问超出其长度的元素将会导致语言运行时异常。与切片的容量相关联的元素只能用于增长切片。在使用这部分元素前，必须将其合并到切片的长度里。

```go
slice :=[] int{10,20,30,40,50}
//创建新切片
newSlice:=slice[1:3]
newSlice[3]=45//对于NewSlice来说并不存在，所以会报错为out of range
```

append：将切片的额外容量合并到切片的长度里

切片增长

切片相对于数组而言的一个好处就是，可以按需增加切片的容量，append函数会处理增加长度时的所有操作细节。

函数append总是会增加新切片的长度，而容量有可能会改变，也可能不会改变，这取决于被操作的切片的可用容量。

```go
newSlice = append(newSlice,60)
```

如果切片的底层数组没有足够的可用容量，append函数会创建一个新的底层数组，将被引用的现有的值复制到新数组里，再追加新的值。如果有容量，那就直接添加

函数append会智能地处理底层数组的容量增长，在切片的容量小于1000个元素时，总是会成倍地增加容量。一旦元素个数超过1000，容量的增长因子会设定为1.25，也就是每次会增加25%的容量。

容量:cap

长度:len

三个索引，可以用来限制容量

```go
slice:=source[2:3:4]//表示起始值，右边界，容量右边界
```

如果有可用容量，会分配一个新的底层数组，，

如果在创建切片时设置切片的容量和长度一样，就可以强制让新切片的第一个append操作创建新的底层数组，与原有的底层数组分离。新切片与原有的底层数组分离后，可以安全地进行后续修改。

append后面可以在一次调用传递多个追加的值。如果使用...运算符，可以将一个切片的所有元素追加到另一个切片里。

```go
s1:=[]int{1,2}
s2:=[]int{3,4}
append(s1,s2...)
Printf(s1)==>1,2,3,4可以全部加进去
```



### :=和=的区别

= 是赋值， := 是声明变量并赋值

### fmt的几种输出区别

```python
Print:   输出到控制台,不接受任何格式化操作
Println: 输出到控制台并换行
Printf : 只可以打印出格式化的字符串。只可以直接输出字符串类型的变量（不可以输出别的类型）
Sprintf：格式化并返回一个字符串而不带任何输出
Fprintf：来格式化并输出到 io.Writers 而不是 os.Stdout
```

当迭代切片的时候，关键字range会返回两个值，第一个值是当前迭代到的索引位置，第二个值是该位置对应元素值的一份副本。

range创建了每个元素的副本，而不是直接返回对该元素的引用。

如果不需要索引值，可以使用占位字符来忽略这个值，占位符:  _

关键字for总是会从切片头部开始迭代，如果想对迭代做更多的控制，依旧可以使用传统的for循环。

```go
slice:=[]int {10,20,30,40}
for index:=2; index<len(slice); index++{
	fmt.Printf("Index: %d Value:%d\n",index,slice[index])
}
```

### 多维切片

可以组合多个切片形成多维切片

```go
slice:=[][]int{{10},{100,200}}
```





在函数间传递切片

在64位的机器上，一个切片需要24字节的内存：指针字段需要8字节，长度和容量字段需要8字节。由于与切片关联的数据包含在底层数组中，不属于切片本身，所以将切片复制到任意函数的时候，对底层数组大小都不会有影响。

复制只会复制切片本身，不会涉及底层数组

**切片相当于（一个指针，其长度和容量），各自占用8个字节，而数组会将整个空间都包含进去，可以理解为数组与指针的结合**，



## 映射

是一个数据结构，用来存储一系列无顺序的键值对

映射里基于键来存储值

映射功能强大之处在于，能够基于键快速减速数据，键就像索引一样，指向与该键关联的值

映射是无序的，意味着没有办法预测键值对被返回的顺序。

映射的实现使用了散列表hash

首先生成散列，进行mod运算，进行映射，并存储，后面只需要再进行对比Hash值就可以找到对应的值

Go语言的映射来说，生成的散列键的一部分，具体来说是低位LOB,被用来选择桶

**映射使用两个数据结构来存储数据。第一个数据结构是数组，内部存储的是用于选择桶的散列键的高八位值，这个数组用于区分每个键值对要存在哪个桶里，第二个数据结构是一个字节数组，用于存储键值对。该字节数组先依次存储了这个桶里所有的键，之后依次存储了这个桶里所有的值。实现这种键值对的存储方式目的在于减少每个桶所需的内存。**

创建和初始化映射

```go
dict:= make(map[string]int)//创建映射，键的类型是string,值的类型是int
```

映射的键可以是任意值。这个值的类型可以是内置的类型，也可以是结构类型，只要这个值可以使用==运算符做比较。

切片、函数以及包含切片的结构类型由于具有引用语义，不能作为映射的键，但是可以作为值使用

```go
dict:=map[[]string]int{}//不可以，不能使用map key type []string
dict:=map[int] []string{}//可以，[]string作为值使用,一个映射键对应一组数据时，非常有用
```

从映射取值有两个选择。

第一个选择是，可以同时获得值，以及一个表示这个值是否存在的标志

```go
value,exits:=colors["Blue"]
if exits{
	fmt.Println(value)
}
```

另一个选择，只返回键对应的值，然后通过判断这个值是不是零值来确定键是否存在

```
value:= colors["Blue"]
if value !=""{
	fmt.Println(value)
}
```

在Go中，通过键来索引映射时，即使这个键不存在也总会返回一个值，返回的是该值对应的类型的零值

映射迭代返回的不是索引和值，而是键值对

```
colors:=map[string]string{
	"AliceBlue": "#f0f8ff",
	"Coral"    : "#ff7F50"
}
for key,value:= colors{
	fmt.Printf(key,value)
}
```

删除时，使用delete函数

```
delete(colors,"Coral")
```

在函数间传递映射不会制造出该映射的一个副本。当传递映射给一个函数，并对这个映射做了修改时，所有对这个映射的引用都会觉察到这个修改

数组是构造切片和映射的基石

切片经常用于处理数据的集合，映射用来处理具有键值对结构的数据

内置函数cap只能用于切片

内置函数len可以用来获取切片或者映射的长度

将切片或者映射传递给函数成本很小，并且不会复制底层的数据结构

 





Go语言中声明用户定义的类型有两种方法，最常用的是struct关键字，可以让用户创建一个结构类型。

```go
type user struct{
	name string
	email string
}
user 就是新类型的名字	
```

使用声明的变量

```go
var bill user 
```

对于bool类型，零值是false

任何时候，创建一个变量并初始化为零值，习惯是使用关键字`var`,这种用法是为了更明确地表示一个变量被设置为零值。

一个短变量声明操作符在一次操作中完成两件事，声明一个变量，初始化变量`:=`

```go
lisa:=user{
	name: "lisa"
	email: "lisa@email.com"
}
```

基于一个已有的类型，将其作为新类型的类型说明。当需要一个可以用已有类型表示的新类型的时候，这个方法会非常好用。

```go
type Duration int64
```

虽然int64是Duration的基础类型，但是Go并不认为Duration和int64是同一种类型，这两个类型是完全不同的有区别的类型。

不能互相赋值，编译器不会对不同类型的值做隐私转换

### 方法

​	方法实际上也是函数，只是在声明时，在关键字func和方法名之间增加了一个参数。

```go
func (u user)notify(){
	fmt.Print()
}
```

`func`和函数名之间的参数叫做接收者，将函数与接收者的类型绑在一起，

如果一个函数有接收者，那么这个函数就被称为**方法**

Go语言中有两种类型的接收者：**值接收者**和**指针接收者**

如果使用值接收者声明方法，调用时会使用这个值的一个副本来执行

使用变量来调用方法

```go
bill.notify()
//与调用一个包里的函数看起来很类似，但是bill不是包名，而是变量名
//在调用notify方法时，使用bill的值作为接收者进行调用，方法notify会接收到bill的值的一个副本
```

也可以使用指针来调用使用值接收者声明的方法

```go
lisa:= &user{"Lisa","lisa@email.com"}
lisa.notify()

可以理解为Go执行如下命令
（*lisa).notify()
```

可以互相调用，

返回值是指针变量的可以用值来接收，返回值是值的也可以用指针变量来接收

### 内置类型

数值类型，字符串类型，布尔类型

当对这些值进行增加或者删除的时候，会创建一个新值

当把这些类型的值传递给方法或者函数时，应该传递一个对应值的副本

### 引用类型

切片，映射，通道，接口和函数类型

当声明上述类型的变量时，创建的变量被称为标头（header）值。

每个引用类型创建的header值是包含一个指向底层数据结构的指针。每个引用类型还包含一组独特的字段，用于管理底层数据结构，因为header是为了复制而设计的，所以永远不需要共享一个引用类型的值。Header值里包含了一个指针，因此通过复制来传递一个引用类型的值的副本，本质上是共享底层数据结构。

## 接口

对接口的调用就是一种多态

接口值并没有具体的行为

### 方法集

| Methods Receivers | Values    |
| ----------------- | --------- |
| (t T)             | T and  *T |
| (t *T)            | *T        |

### 结构体

```
type Feed struct{
	Name string `json:"site"`
    URI  string `json:"link"`
    Type string `json:"type"`
}
```

``里面的部分被称为标记（tag)

标记描述了JSON解码的元数据，用于创建Feed类型值的切片。每个标记将结构类型里字段对应到JSON文档里指定名字的字段。

内部类型的标识符提升到了外部类型



当一个标识符的名字以**小写字母开头**时，这个标识符就是**未公开**的，即包外的代码不可见，如果一个标识符以**大写字母开头**，这个标识符就是**公开的**。

## 并发

WaitGroup是一个计数信号量，可以用来记录并维护运行的goroutine,如果WaitGroup的值大于0，Wait方法就会阻塞。

为了减少WaitGroup的值并最终释放main函数，使用defer声明在函数退出时调用Done方法





对一个共享资源的读和写必须是原子化的，同一时刻只能有一个goroutine对共享资源进行读和写操作。

如果两个或者多个goroutine在没有相互同步的情况下，访问某个共享的资源，并试图同时读和写这个资源，就处于相互竞争的状态，被称作**竞争状态**

一种修正代码、消除竞争状态的办法是，使用Go语言提供的锁机制，来锁住共享资源，从而保证goroutine的同步状态。

AddInt64,LoadInt64,StoreInt64,都是atomic的方法，可以强制同一时刻只有一个goroutine运行并完成操作

第二种方法是互斥锁（mutex)，用于在代码上创建一个临界区，保证同一时间只有一个goroutine可以执行这个临界区代码

```go
mutex sync.Mutex
mutex.Lock()
mutex.Unlock()
```

可以使用通道，通过发送和接收需要共享的资源，在goroutine之间做同步

当一个资源需要在goroutine之间共享时，通道在goroutine之间架一个通道，并提供了确保同步交换数据的机制。

创建通道

```go
unbuffered:= make(chan int )//无缓冲的整型通道
buffered := make(chan string ,10)//有缓冲的字符串通道
```

make的第一个参数需要是关键字chan,之后跟着允许通道交换的数据的类型。

如果创建的是一个有缓冲的通道，还需要在第二个参数指定这个通道的缓冲区的大小

向通道发送值或者指针需要用到<-操作符

```go 
buffered := make(chan string ,10)//有缓冲的字符串通道
buffered <-"Gopher"//通过通道发送一个字符串
```

从通道中接收值

```go
value:= <-buffered
```

当从通道里接收一个值或者指针时，<-运算符在要操作的通道变量的左侧

### 无缓冲通道

指在接收前没有能力保存任何值的通道。这种类型的通道要求发送ｇｏｒｏｕｔｉｎｅ和接收同时准备好，才能完成发送和接收操作。如果两个ｇｏｒｏｕｔｉｎｅ没有同时准备好，通道会导致先执行发送或接收操作的ｇｏｒｏｕｔｉｎｅ阻塞等待。这种对通道进行发送和接收的交互本身就是同步的。

调用Ｄｏｎｅ（），ＷａｉｔＧｒｏｕｐ就会减一

### 有缓冲通道

只有在通道中没有要接收的值时，接收动作才会阻塞。只有在通道没有可用缓冲区容纳被发送的值时，发送才会被阻塞。无缓冲的通道保证进行发送和接收的ｇｏｒｏｕｔｉｎｅ会在同一时间进行数据交换，有缓冲的通道没有这个保证



关闭通道close()



```
time.Duration()//Duration 类型用于表示两个时刻 ( Time ) 之间经过的时间，以 纳秒 ( ns ) 为单位。
```

