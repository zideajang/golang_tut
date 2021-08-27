## defer 调用
在介绍 defer 之前，我们先看一个例子做一个小实验。很多现代的编程语言中都有 `defer` 关键字，Go 语言的 `defer` 会在当前函数返回前执行传入的函数

```go
func main() {
	fmt.Println("start")
	fmt.Println("middle")
	fmt.Println("end")
}
```
```
start
middle
end
```
在控制台输出顺序和程序执行顺序没有任何差别，当将其中 `fmt.Println("middle")` 前面添加 `defer` 关键字
```go
func main() {
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
}
```
```
start
end
middle
```
大家可以已经发现了 middle 输出最后，这是因为 defer 将会在其所在的函数退出或返回前才执行，当 main 函数要退出时才执行 ```defer fmt.Println("middle")```。
```go
func main() {
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
}
```
```go
end
middle
start
```
如果这里有多个 defer 那么他们执行顺序又会怎样呢？他们执行顺序会遵循 LIFO 也就是最后一个会被最先执行的原则。
为什么将 defer 执行的顺序这样设计呢？这是因为 defer 我们通常都是用来关闭资源，所以采用这种先进后出方式执行 defer 是有一定道理的。

```go
func main() {
	res, err := http.Get("https://xxx/xxx.txt")
	if err != nil{
		log.Fatal(err)
	}

	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s",robots)
}
```
多数情况我们需要打开资源一段时时间，然后进行操作。然而这样一来我们就很容易忘记关闭资源。这时候我们可以同 defer 来定义在退出函数时将这些打开的资源一一关闭
```defer res.Body.Close()```

有的时候我们需要程序发生错误后需要进行后续的处理，例如关闭文件、结束网络连接等。这时候我们就会用到 def ,如果是 java 的 developer 可以理解为 finally。

- 确保在函数结束时发生
- 参数在 defer 语句时计算
- defer 列表为后进先出

```go
var wg sync.WaitGroup

func say(s string){
	for i:=0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}

	wg.Done()
}

func main(){
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("There")

	wg.Add(1)
	go say("Hi")
	wg.Wait()
}
```

```go

func say(s string){
    defer wg.Done()
	for i:=0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}

}
```


```go
func tryDefer(){
	fmt.Println(1)
	fmt.Println(2)
}

func main(){
	tryDefer()
}
```
```
1
2
```

```go
func tryDefer(){
	defer fmt.Println(1)
	fmt.Println(2)
}

```

使用 defer 就会
```
2
1
```
defer 相对来说是一个栈先进后出
```go
func tryDefer(){
	defer fmt.Println(1)
	defer fmt.Println(2)
    fmt.Println(3)
    
}
```

```
3
2
1
```

加了 defer 就不怕程序中间 return 甚至是 panic

```
func tryDefer(){
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
	fmt.Println(4)
}
```


```go
func writeFile(filename string){
	file,err := os.Create(filename)
	if err !=nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()

	for i := 0; i < 20; i++{
		fmt.Fprintln(writer,f())
	}
}
```

避免我们忘记，可以只写 defer 

```go
for i:=0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
    }
```

```go
3
2
1
0
panic: printed too many
```