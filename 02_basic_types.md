默认了解过一门语言，也就是在开始学习 golang 默认已经

- 如何定义一个变量
- 定义变量是是否需要指定类型
- 定义变量是否需要初始化赋值

## golang 中基础类型

```golang
var name string = "machine leanring golang"
```

```golang
var name string
name = "machine leanring golang"
```

### 整型(Integers)
#### 无符号整型(unsigned integers)
- uint8
- uint16
- uint32
- uint64

#### 符号整型(unsigned integers)
- int8
- int16
- int32
- int64

#### Machine Dependent type
根据机器系统是 32 位还是 64 位来确定 int 使用 int64 还是 int32
- uint
- int
- uintptr

#### Complex
除非我们在写物理引擎或者一些计算，我们几乎是用不到 Complex 这个数据类型 


### 变量赋值
- golang 在赋值时候无需显式指定变量类型
- 推荐类型有时候会存在风险，推荐类型，随后给一个比较大的值

```golang
var number = 20
```

#### 输出类型

```golang
var number = 20
fmt.Printf("%T",number) //int
```
使用 `fmt` 包提供的 `Printf` 方法时使用 `%T` 通配符就可以提取输入变量的类型



## slice

go 语言中给我们带来了许多新的东西，其中就包括这个 `slice` 那么 `slice` 翻译为切片又是什么东西，有用什么不同，需要我们怎么理解 slice 和 array 的不同呢，带着这些疑问开始今天的分享。


### slice 创建

在 golang 语言中，定义一个数组 arr 方式如下

```go
arr := [3]int{1,2,3}
```

而定义一个 slice 方式为，大家可能发现了在定义 slice 不需要给出长度。

可以将 slice 理解为数组的引用。

```go
s := []int{1,2,3}
```

下面创建 slice 需要大家注意一下，前 3 个元素为 0, 1, 2, 而 `5: 100` 表示索引为 5 也就是第 6 元素 

```go
s2 := []int{0, 1, 2, 5: 100} // [0 1 2 0 0 100] 6 6
```



```cpp
struct slice{
	byte* array;
	uintgo len;
	uintgo cap;
}
```



在 golang 这门语言中，可以是使用 `make` 语句创建只有 `slice` `map` 和 `channel` 所以也可以通过 `make`来创建 slice。



```go
	s3 := make([]int, 3, 3)
	s5 := make([]int, 3, 6)

	fmt.Println(s3, len(s3), cap(s3))
	fmt.Println(s5, len(s5), cap(s5))
```



```go
[0 0 0] 3 3
[0 0 0] 3 6
```



```go
	arr := [...]int{0,1,2,3,4,5,6}
	s6 := arr[1:4:5]
```

