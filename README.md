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
        c = iota //b=1)
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
-`goroutine-channel 中，无缓冲channel数据流入流出顺序，是先入先出的`

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
##延迟列表
```go
for i:=0; i<5; i++ {
defer fmt.Printf("%d ", i)
}
延迟函数是按照后进先出的顺序执行，所以上面这段代码的执行结果是43210
```

## To Be Continue...

## Ref:
1. [http://blog.csdn.net/kjfcpua/article/details/18265441](http://blog.csdn.net/kjfcpua/article/details/18265441)  
2. [https://blog.csdn.net/u012291393/article/details/78378493](https://blog.csdn.net/u012291393/article/details/78378493)
3. 学习go语言.pdf
