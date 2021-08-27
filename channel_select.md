 并发(concurrency) 更多的是一种设计(design) 



#### CSP(Communication Sequential Processes)

- Tony Hoare 1978 年提出的 CSP
- Each process is bult for sequential execution
- Data is communicated between processes via channels No shared state
- Scale by adding more of the same



#### Go 并发的工具集(toolset)

- go routines

- channels

- select 

- sync package









#### 什么是 channel

channel 可以看做 Goroutines 间用来通信的管道，在 go 语言中，两个 Goroutines 就是通过 channel 来交互数据达到同步数据，避免冲突。



#### 如何创建 channel

在定义 channel 时需要给出一个类型，和 cpp 的指针有点类似，估计都是开辟一块内存，为 channel 指定了类型之后，该 channel 就只能接受指定类型的数据，不能接受其他类型的数据，在输出通道类型时就是就是你指定类型

Channel 的初始值是 `nil`。`nil` 通道没有任何用处，需要用 `make(chan int)` 



> make 这函数是一个内建函数，其第一个参数是类型，第二个参数是长度。Go 语言中初始化一个结构时会用到 make 和 new 都是初始化一个结构体，返回一个结构体的指针，但是 make 要相对 new 要复杂一些

```go
package main

import "fmt"

func main() {  
    var a chan int
    if a == nil {
        fmt.Println("channel a is nil, going to define it")
        a = make(chan int)
        fmt.Printf("Type of a is %T", a)
    }
}
```



```
channel a is nil, going to define it  
Type of a is chan int  
```



#### 通过 channel 发送和接受数据

```go
data := <- a // read from channel a  
a <- data // write to channel a  
```



通过箭头相对于通道的方向来表示通道发送数据还是接收数据。在第一行中，箭头从 a 向外指向 data，表示从通道 a 中读取数据并将其存储到变量 data 中。在第二行中，箭头指向 a，表示正在向通道 a 写入数据。



#### 发送和接收默认为阻塞
channel 的发送和接收默认是阻塞的，这意味着什么？当数据被发送到一个 channel 时，控制在发送语句中被阻断，直到其他 Goroutine从该通道读取。同样地，当数据从一个 channel 中读出时，读被阻塞，直到某个 Goroutine 将数据写到该  channel。也就是创建好一个通道，无论是先写入还是先读取通道都会让发送和接受(写入和读取)的所在 Goroutine 发生阻塞。



This property of channels is what helps Goroutines communicate effectively without the use of explicit locks or conditional variables that are quite common in other programming languages.

channel 的这一属性有助于 Goroutine 有效地进行通信，而无需使用显式锁或条件变量，这在其他编程语言中是很常见的。



上面有关如何声明一个 channel 以及如何通过 channel 在 goroutine 间发送和接受数据

```go
package main

import (  
    "fmt"
    "time"
)

func sayhello() {  
    fmt.Println("Hello  goroutine")
}
func main() {  
    go sayhello()
    time.Sleep(1 * time.Second)
    fmt.Println("main function")
}
```



这里我们在这里使用了一个 `sleep` ，让主 Goroutine 等待sayhello Goroutine 的结束。这个简单休眠就是为了让 sayhello 这个 Goroutine 执行完成后在退出程序

```go
package main

import (  
    "fmt"
)

func sayHello(done chan bool) {  
    fmt.Println("Hello goroutine")
    done <- true
}
func main() {  
    done := make(chan bool)
    go sayHello(done)
    <-done
    fmt.Println("main function")
}
```



通过通道方式来实现 `time.Sleep` 效果等待 `sayHello` 执行完毕再去退出程序，这样通过休眠固定时间显然是不合理的。上面介绍过了有关 channel 可以起到一定阻塞的作用替换掉 `time.sleep` 的效果。



现在，我们的 主Goroutine 将被阻塞，直到 done 通道上的数据。sayHello Goroutine 接收这个通道作为参数，在 sayHello Goroutine 内部打印输出` Hello  goroutine`，然后写到 done channel。当这个写入完成后，main Goroutine从doed 通道接收数据，这是主 Goroutine 也就解除阻塞，然后打印输出`main  funciton` 。



在使用 channel 时，需要注意一个重要问题就是死锁 (deadlock)。如果一个 Goroutine 在一个 channel 上发送数据，那么预期其他的Goroutine 应该会接收这个写入到 channel 的数据。如果没有其他 Goroutine 去接受数据，那么程序在运行时就会出现死锁(deadlock)的恐慌。



类似的，如果一个 Goroutine 正在等待从一个 channel 接收数据，那么其他一些 Goroutine 就应该在这个通道上写数据，否则程序就会恐慌。



```go
package main

func main() {  
    ch := make(chan int)
    ch <- 5
}
```



### 关闭通道和通道上的范围循环



发送者有能力关闭 channel，以通知接收者不再向该 channel 上发送数据。接收者在从 channel 接收数据时，可以使用一个的变量来检查 channel 是否已经关闭。

```go
v, ok := <- ch  
```



在上面的语句中，如果 channel 还没有关闭，那么 ok 就是 true。如果 ok 是 false 就这意味着当前要读取数据的 channel 以及关闭。从一个关闭的通道中读取的值将是该通道类型的初始值。例如，如果通道是一个 int channel，那么从一个关闭的通道收到的值将是 0。

```go
package main

import (  
    "fmt"
)

func producer(chnl chan int) {  
    for i := 0; i < 10; i++ {
        chnl <- i
    }
    close(chnl)
}
func main() {  
    ch := make(chan int)
    go producer(ch)
    for {
        v, ok := <-ch
        if ok == false {
            break
        }
        fmt.Println("Received ", v, ok)
    }
}
```



在上面的程序中，生产者 Goroutine 将 0 到 9 写入 `chnl` 通道，然后关闭该通道。main 函数在有一个无限的 for 循环，其中用变量 `ok`检查  channel 是否被关闭。 如果 `ok` 是` false`，意味着 channel 已经关闭，循环将被中断。否则，接收到的值和 `ok` 的值被打印出来。



```go
Received  0 true  
Received  1 true  
Received  2 true  
Received  3 true  
Received  4 true  
Received  5 true  
Received  6 true  
Received  7 true  
Received  8 true  
Received  9 true 
```



`for range`形式的`for`循环可以用来从一个通道接收数值，直到它被关闭。

让我们用for range循环重写上面的程序。





## channel 和 select 

- Channel 的结构设计
- 



#### 平行计算

不同人操作(更新)同一个数据，目的是为了解决多线程的一致性。



### 共享内存

有一块共享内存，大家同时操作一块共享内容，同一时刻只允许一个人操作这块内存。



### 消息通信(拷贝内存)

Tony Hoare 1977 年基于消息通信使用 channel 原理提出并发的 Communication Sequential Processes(CSP) 数学原理。设计参考 CSP 都是由不，也有共享



### buffer 和 unbuffer

阻塞式 channel 和非阻塞的 channel，对于 **buffered channel** 产生数据并存入 buffer，然后 reader 从 buffer 中进行消费，而 **unbuffered channel** 会阻塞到 reader 从 channel 中读取数据。



- zero-case 会发生永久阻塞



### golang 调度器

