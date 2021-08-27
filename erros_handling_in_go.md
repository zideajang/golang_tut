

> 我们通常会有做错，什么是错误，错误又是相对于谁的，是相对于自己的，还是性对于别是错误。出现错误后应该如何处理，小朋友弄坏电器，找爸爸妈妈来就休息。



与其他主流语言如 Javascript、Java 和 Python 相比，Golang 的错误处理方式可能和这些你熟悉的语言有所不同。所以才有了这个想法根大家聊一聊 golang 的错误处理方式，以及实际开发中应该如何对错误进行处理。因为分享面对 Golang有一个基本的了解 developers, 所以一些简单地方就不做赘述了。



### 如何定义错误

在 golang 语言中，无论是在类型检查还是编译过程中，都是将错误看做值来对待，和 string 或者 integer 这些类型值并不差别。声明一个 string 类型变量和声明一个 error 类型变量是没什么区别的。

你可以定义接口作为 error 的类型，有关 error 能够提供什么样信息都是由自己决定的，这是 error 在 golang 作为值的好处，不过这样做也自然有其坏处，有关 error 定义好坏就全由其定义开发人员所决定，也就是有关 error 融入过多人为的主观因素。



```go
package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	dir, err := ioutil.TempDir("","temp")

	if err != nil{
		fmt.Errorf("failed to create temp dir: %v",err)
	}
}
```





#### 错误在语言中的重点地位

在 Go 语言中错误处理设计一直大家喜欢讨论的内容，错误处理是该语言的核心，但该语言并没有规定如何处理错误。社区已经为改进和规范错误处理做出了努力，但许多人忽略了错误在我们应用程序领域中的核心地位。也就是说，错误与客户和订单类型一样重要。

### Golang中的错误



错误表示在应用程序中发生了不需要的情况。比方说，想创建一个临时目录，在那里可以为应用程序存储一些文件，但这个目录的创建失败了。这是一个不期望的情况，就可以用错误来表示。



通过创建自定义错误可以将更丰富错误信息传递给调用者。个返回值返回将错误交给调用函数人来处理错误。Golang 本身允许函数具有多个返回值，所以通常把错误作为函数最后一个参数返回给调用者来处理。



#### errors 是 I/O

- 有时候开发人员是 error 的生产者(写 error)
- 有时候开发人员又是 error 的消费者(读 error)

也就是我们开发程序一部分工作是读取和写入 error



#### errors 的上下文

什么是 error 的上下文呢? 如何定义 error 需要考虑一些因素，例如在不同程序我们定义 error 和处理 error 方式也不仅相同

- CLI 工具
- 库
- 长时间运行的系统

而且我们需要考虑使用程序的人群，他们是什么方式来使用系统，这些因素都是我们设计也好定义错误信息要考虑的因素。



#### 错误的类型

就错误核心，那么错误可能是我们预料之中的错误，错误也可能是我们没有考虑到，例如无效内存，数组越界，也就是单靠代码自身暂时是解决不了的错误 ，这样的误差往往让代码恐慌，所以 Panic。通常这样错误对于程序是灾难性的失败，无法修复的。





### 自定义错误

如前所述，错误使用内置的错误接口类型来表示，其定义如下。

```go
type error interface {  
    Error() string
}
```



下面举了 2 例子来定义 `error` ，分别定义两个 `struct` 都实现了 `Error()` 接口即可



```go
type SyntaxError struct {
    Line int
    Col  int
}

func (e *SyntaxError) Error() string {
    return fmt.Sprintf("%d:%d: syntax error", e.Line, e.Col)
}
```



```golang
type InternalError struct {
    Path string
}

func (e *InternalError) Error() string {
    return fmt.Sprintf("parse %v: internal error", e.Path)
}
```





该接口包含一个方法`Error()`，以字符串形式返回错误信息。每一个实现了错误接口的类型都可以作为一个错误使用。当使用` fmt.Println` 等方法打印错误时，Golang 会自动调用` Error() `方法。

在 Golang 中，有多种创建自定义错误信息的方法，每一种都有自己的优点和缺点。



#### 基于字符串的错误
基于字符串的错误可以用 Golang 中两个开箱即用方法来自定义错误，适用哪些仅返回描述错误信息的相对来说比较简单的错误。



```go
err := errors.New("math: divided by zero")
```



将错误信息传入到`errors.New()`方法可以用来新建一个错误



```go
err2 := fmt.Errorf("math: %g cannot be divided by zero", x)
```



`fmt.Errorf` 通过字符串格式方式，可以将错误信息包含你错误信息中。也就是为错误信息添加了一些格式化的功能。



#### 自定义数据结构的错误

可以通过在你的结构上实现`Error`接口中定义的`Error()`函数来创建自定义的错误类型。下面是一个例子。



### Defer, panic 和 recover

Go 并不像许多其他编程语言（包括 Java 和 Javascript ）那样有异常，但有一个类似的机制，即 "Defer, panic 和 recover"。然而，panic 和 recover 的使用情况与其他编程语言中的异常非常不同，因为代码本身无法应对时候和不可恢复的情况下使用。

### Defer

有点类似析构函数，在函数执行完毕后做一些资源释放等收尾工作，好处其执行和其在代码中位置并没有关系，所以可以将其写在你读写资源语句后面，以免随后忘记做一些资源释放的工作。关于 `defer` 输出也是面试时，面试官喜欢问的一个问题。



```go
package main

import(
	"fmt"
	"os"
)

func main(){
	f := createFile("tmp/machinelearning.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil{
		panic(err)
	}
	return f
}

func closeFile(f *os.File){
	fmt.Println("closing")
	err := f.Close()

	if err != nil{
		fmt.Fprintf(os.Stderr, "error:%v\n",err)
		os.Exit(1)
	}
}

func writeFile(f *os.File){
	fmt.Println("writing")
	fmt.Fprintln(f,"machine leanring")
}
```



`defer` 语句会将函数推入到一个栈结构中。同时栈结构中的函数会在 `return`语句执行后被调用。



```go
package main


import "fmt"

func main(){
	// defer fmt.Println("word")
	// fmt.Println("hello")

	fmt.Println("hello")
	for i := 0; i <=3; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("world")
}
```



```go
hello
world
3
2
1
0
```

可以通过在你的结构上实现`Error`接口中定义的`Error()`函数来实现自定义错误类型，下面是一个例子。

### Panic

panic 语句向 Golang 发出信号，这时通常是代码无法解决当前的问题，所以停止代码的正常执行流程。一旦调用了 panic，所有的延迟函数都会被执行，并且程序会崩溃，其日志信息包括 `panic`值（通常是错误信息）和堆栈跟踪。

举个例子，当一个数字被除以0时，Golang会出现 panic。

```go
package main

import "fmt"

func main(){
	divide(5)
}

func divide(x int){
	fmt.Printf("divide(%d)\n",x+0/x)
	divide(x-1)
}
```



```go
divide(5)
divide(4)
divide(3)
divide(2)
divide(1)
panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.divide(0x0)
        /Users/zidea2020/Desktop/mysite/go_tut/main.go:10 +0xdb
main.divide(0x1)
        /Users/zidea2020/Desktop/mysite/go_tut/main.go:11 +0xcc
main.divide(0x2)
        /Users/zidea2020/Desktop/mysite/go_tut/main.go:11 +0xcc
main.divide(0x3)
        /Users/zidea2020/Desktop/mysite/go_tut/main.go:11 +0xcc
main.divide(0x4)
        /Users/zidea2020/Desktop/mysite/go_tut/main.go:11 +0xcc
main.divide(0x5)
        /Users/zidea2020/Desktop/mysite/go_tut/main.go:11 +0xcc
main.main()
        /Users/zidea2020/Desktop/mysite/go_tut/main.go:6 +0x2a
exit status 2
```





### Recover

Go语言提供了recover内置函数，前面提到，一旦panic，逻辑就会走到defer那，那我们就在defer那等着，调用recover函数将会捕获到当前的panic，被捕获到的panic就不会向上传递了。然后，恢复将结束当前的 Panic 状态，并返回 Panic 的错误值。

```go
package main

import "fmt"

func main(){
	accessSlice([]int{1,2,5,6,7,8}, 0)
}

func accessSlice(slice []int, index int) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("internal error: %v", p)
		}
	}()

	fmt.Printf("item %d, value %d \n", index, slice[index])
	defer fmt.Printf("defer %d \n", index)
	accessSlice(slice, index+1)
}
```



### 包装错误

Golang 也允许对错误进行包裹，通过错误嵌套，在原有错误信息上添加一个额外信息帮助调用者对问题判断以及后续应该如何处理信息。以通过使用`%w`标志和`fmt.Errorf`函数来对原有的错误进行保存提供一些特定的信息，如下例所示。

```go
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	err := openFile("non-existing")

	if err != nil {
		fmt.Printf("error running program: %s \n", err.Error())
	}
}

func openFile(filename string) error {
	if _, err := os.Open(filename); err != nil {
		return fmt.Errorf("error opening %s: %w", filename, err)
	}

	return nil
}
```



上面已经通过代码演示如何包装一个错误，程序会打印输出使用`fmt.Errorf`添加文件名的包装过的错误，也打印了传递给`%w`标志的原有错误信息。这里再补充一个 Golang 还提供的功能，通过使用`error.Unwrap`来还原错误信息，从而获得原有的错误信息。

```go
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	err := openFile("non-existing")

	if err != nil {
		fmt.Printf("error running program: %s \n", err.Error())

		// Unwrap error
		unwrappedErr := errors.Unwrap(err)
		fmt.Printf("unwrapped error: %v \n", unwrappedErr)
	}
}

func openFile(filename string) error {
	if _, err := os.Open(filename); err != nil {
		return fmt.Errorf("error opening %s: %w", filename, err)
	}

	return nil
}
```



### 错误的类型转换

有时候需要在不同的错误类型之间进行转换，有情况需要通过类型转换来为错误添加信息，或者换一种表达方式，。`errors.As`函数提供了一个简单而安全的方法，通过寻找错误链中匹配错误类型进行转化输出。如果没有找到匹配的，该函数返回`false`。



```go
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main(){
	// Casting error
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *os.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
}
```



在这里，试图将通用错误类型转换为`os.PathError`，这样就可以访问该特定的错误信息，这些信息保存在结构体中的` Path` 属性上。

#### 错误类型检查

Golang 提供了`errors.Is`函数来用于检查错误类型是否为指定的错误类型，该函数返回一个布尔值值来表示是否为指定错误类型。



```go
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main(){
	// Check if error is a specific type
	if _, err := os.Open("non-existing"); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("file does not exist")
		} else {
			fmt.Println(err)
		}
	}
}
```

