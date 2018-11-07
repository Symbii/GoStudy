# GoStudy
- `package name,name一般是目录名字,一个工程一定会有一个package main,
  import "name" name 同样指的是目录名字,代码中才是包的名字`
- `GOROOT 指的是go目录, GOPATH指的是src的目录`
- `任何赋给_的值都将被丢弃`
- `安装好Go语言之后，其文档可以通过go doc命令查看，例如查看fmt，在命令行输入go doc fmt即可。如果要查看某个包中的子目录的文档，使用go doc parent_package/son_directory,例如go doc hash/fnv`
- `iota 只能在const类型中使用，每当const关键字出现的时候就会被重置`
```go
    const a = iota //a=0   
    const (
        b = iota  //b=0
		c = iota  //c=1
	)
```
-   `自定义类型,自增长常量经常包含一个自定义枚举类型，允许你依靠编译器完成自增设置。`
```go
        type Stereotype int         
          const ( s
              TypicalNoob Stereotype = iota // 0 
              TypicalHipster                // 1 
              TypicalUnixWizard             // 2 
              TypicalStartupFounder         // 3 
          ) 
```

- `跳过某些值`
```go
        type Stereotype int         
          const ( 
              TypicalNoob Stereotype = iota // 0 
              _                             // 1 
              _                             // 2 
              TypicalStartupFounder         // 3 
          ) 
```
-  `goroutine-channel:无缓冲channel和缓冲channel，实现并发。无缓冲channel，空的时候不停地取或者写入过未取不停地写会导致死锁，但是会有一种情况：main又没等待其它goroutine，自己先跑完了`
```go
func main() {    
    c := make(chan int)
    go func() {
       c <- 1
    }()
}
```
- `goroutine-channel 中，无缓冲channel数据流入流出顺序，是先入先出的`

```go
package main

import "fmt"

var ch chan int = make(chan int)

func foo(id int) { //id: 这个routine的标号
	fmt.Printf("push %d \n", id)
	ch <- id
}

func main() {
	// 开启5个routine
	for i := 0; i < 5; i++ {
		go foo(i)
	}

	// 取出信道中的数据
	for i := 0; i < 5; i++ {
		fmt.Printf("pull %d \n", <- ch)
	}
}
按照上面的代码运行，最先写入的数据并不固定，但是数据存取顺序按照先入先出
```
## 延迟列表
```go
for i:=0; i<5; i++ {
defer fmt.Printf("%d ", i)
延迟函数是按照后进先出的顺序执行，所以上面这段代码的执行结果是43210
func test()(ret int){
	 defer func() {
		ret++
	}()
	return 1
}
延迟函数是在函数结束前执行，所以上面test 返回值是2
```

```go
func test1(arg ...int) (ret []int){
	if 1 == len(arg){
		return arg
	}
	test1(arg[:2]...)
	return test1(arg...)
}
变参传递可以原样传递 ，也可以切片传递部分, arg 是一个slice type.
```

## 回调函数
````go
func callback(y int, f func(int)) {
	f(y)  //调用回调函数 f 输入变量 y
}
````

## Panic&Recover
````go
 func throwsPanic(f func()) (b bool) { 
 	defer func() { 
        if x := recover(); x != nil { 
 	    b = true
        } 
    }()
  f() 
  return 
}
recover只在延迟函数中有效，检测f函数在运行中 是否产生panic
````

## Method
```
method 必须要有一个recieve
func (s structname) funcname(...)(...){
}
这样子之后funcname这个函数就与s绑定,可以通过s.funcname直接访问，s就是funcname的接收者
```
## Map

```go
package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

func main() {
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
        40.68433, 74.39967,
    }
    m["test"] = Vertex{
        12.0, 100,
    }
    fmt.Println(m["Bell Labs"])
    fmt.Println(m)
}
打印的结果，map的顺序是随机的,如果key在map中存在，则访问map会返回对应的值，如果不存在则会返回两个值第二个值为bool类型，第一个值为对应value类型的0值
age, ok = map["haha"]
if !ok--->if haha is not key ， age == 0
```
[some question about map order](https://stackoverflow.com/questions/11853396/google-go-lang-assignment-order)

## http.Get

```go
package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main(){
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b , err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil{
			fmt.Fprintf(os.Stderr, "fetch:  reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
http.Get(url) 遇到一个错误（fetch: Get www.baidu.com: unsupported protocol scheme ""），因为必须url必须输入为https(http):\\....
```

## Something about Paxos
### 假设有多个proposer都想要向整个网络环境发送一条命令，让整个环境都执行，如何达到一致

>初始化

| Proposer    |      Acceptor      |
|-------------|--------------------|
| c:等待执行命令 <br> t=0 当前尝试的票号，票可以理解为为了区分时间前后|  T<sub>max</sub> = 0 当前已发布的最大票号，即时间最新<br>C = X  当前储存的命令<br>T<sub>store</sub>= 0 与当前存储的命令C对应的票号 | 

>一阶段：
 
| Proposer    |      Acceptor      |
|-------------|--------------------|
|1: t = t+1 <br> 2:向所有Acceptor发送消息，请求得到编号为目前t的票即t=t+1|  3：if t>T<sub>max</sub> 就说明这次请求是最新的，那么Acceptor 就使T<sub>max</sub>=t，并且回复给proposer 目前存储的命令和相对应的票号（T<sub>store</sub>， C）| 

>二阶段

| Proposer    |      Acceptor      |
|-------------|--------------------|
|4:如果过半数服务器回复ok：<br> 选择T<sub>store</sub>值最大的（T<sub>store</sub>，C）<br> 5:如果T<sub>store</sub>>0 then c=C <br> 6:向这些回复了ok的acceptor发送消息：propose(t,c) <br> Caution:5和6都是基于4的基础上进行，如果4不满足则从新开始1|  7：如果t=T<sub>max</sub>，那么C=c,T<sub>store</sub>=t，并且回复：Success <br> Caution:这里t=Tmax不是一定成立的，虽然在一阶段的时候已<br>经将T<sub>max</sub>更新为t，但是整个过程中还有其他Proposer会更新<br>Acceptor的T<sub>max</sub>，一旦别的proposer更新T<sub>max</sub>，将导致当前的<br>Proposer重新回到一阶段| 

>三阶段

| Proposer    |      Acceptor      |
|-------------|--------------------|
| 8：如果过半数服务器回复success then，向每个服务器发送execute(c)| |


## Go语言局部变量内存分配
> go语言会按照变量在函数结束的时候是否还有引用，自动选择这个这个局部变量是在堆还是栈上分配。var 或者 new 一个变量并不能决定
```go
var global *int

func fool(){
	var x int
	global = &x
	return
}

func  fool1(){
	y:= new(int)
	*y = 1
	return
}

在fool中x变量会在堆上分配，在fool1中y会在栈上分配。
局部变量的生命周期，只取决于是否地址还有引用或者说还可以被其他变量访问，包括在一个循环体中定义的变量，但是生命周期和局部作用域不是一个含义，局部变量的生命周期出现上面描述的情况会长于作用域

```

## Slice
```go
var s []int
s = nil
s = []int(nil)
s = []int{}---->len(s) == 0 & s != nil,个人理解是此时 s已经指向了底层生成的一个数组，只不过slice长度为0 ，所以判断一个slice是否时空，应该使用len(s)== 0 判断
slice[m:n] n不能大于cap（slice），大于的话会触发一个panic，但是可以大于len（slice）这样子会使得slice增长

slice 复制只是增加了一个别名，slice只能同 nil比较如上。

内置的append函数有一套内存扩展策略，因此我们并不能确认新的slice和原始的slice是否引用的是相同的底层数组空间。同样，我们不能确认在原先的slice上的操作是否会影响到新的slice。所以将append返回值直接赋值给输入的dst是很有必要的。
runes = append(runes, r)
```

## To Be Continue...
```
工作很忙，挤挤还是有的。
人还是要有梦想，万一梦想没实现，发财了呢？
```
## Ref:
1. [http://blog.csdn.net/kjfcpua/article/details/18265441](http://blog.csdn.net/kjfcpua/article/details/18265441)  
2. [https://blog.csdn.net/u012291393/article/details/78378493](https://blog.csdn.net/u012291393/article/details/78378493)
3. 学习go语言.pdf
