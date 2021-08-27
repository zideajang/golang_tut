默认了解过一门语言，也就是在开始学习 golang 默认已经

- 如何定义一个变量
- 定义变量是是否需要指定类型
- 定义变量是否需要初始化赋值

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