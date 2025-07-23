## 练习 7.7： 解释为什么帮助信息在它的默认值是20.0没有包含°C的情况下输出了°C。

### 了解 `flag.Duration` 的实现
先了解 flag包的实现，以 `flag.Duration("period", 1*time.Second, "sleep period")` 举例

```go
// Duration defines a time.Duration flag with specified name, default value, and usage string. Duration定义一个时间。具有指定名称、默认值和使用字符串的持续时间标志。
// The return value is the address of a time.Duration variable that stores the value of the flag. 返回值是一个时间的地址。存储标志的值的持续时间变量。
// The flag accepts a value acceptable to time.ParseDuration. 标志接受时间可接受的值。ParseDuration。
func Duration(name string, value time.Duration, usage string) *time.Duration {
	return CommandLine.Duration(name, value, usage)
}
// 本质是封装 CommandLine 的 Duration 方法
```

```go
// Duration defines a time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationVar(p, name, value, usage)
	return p
}
```

```go
// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	f.Var(newDurationValue(value, p), name, usage)
}
// 这个 f.Var 就是 （7.4 tempconv）里 flag.CommandLine.Var(&f, name, usage) 注册自定义标志的方法
```

```go
// 生成符合flag.Value接口的time.Duration
// -- time.Duration Value
type durationValue time.Duration

func newDurationValue(val time.Duration, p *time.Duration) *durationValue {
	*p = val
	return (*durationValue)(p)
}

func (d *durationValue) Set(s string) error {
	v, err := time.ParseDuration(s)
	if err != nil {
		err = errParse
	}
	*d = durationValue(v)
	return err
}

func (d *durationValue) String() string { return (*time.Duration)(d).String() }

//...
```

了解该实现，即对我们自定义标志有帮助，同时也意识到之前一定要用指针作为返回值的原因

### 什么时候执行 Set ，即根据标志更新数据
`flag.Parse()`

```go
// Parse parses the command-line flags from [os.Args][1:]. Must be called
// after all flags are defined and before flags are accessed by the program.
func Parse() {
	// Ignore errors; CommandLine is set for ExitOnError.
	CommandLine.Parse(os.Args[1:])
}
```

```go
// Parse parses flag definitions from the argument list, which should not
// include the command name. Must be called after all flags in the [FlagSet]
// are defined and before flags are accessed by the program.
// The return value will be [ErrHelp] if -help or -h were set but not defined.
func (f *FlagSet) Parse(arguments []string) error {
	f.parsed = true
	f.args = arguments
	for {
		seen, err := f.parseOne()
		if seen {
			continue
		}
		if err == nil {
			break
		}
		switch f.errorHandling {
		case ContinueOnError:
			return err
		case ExitOnError:
			if err == ErrHelp {
				os.Exit(0)
			}
			os.Exit(2)
		case PanicOnError:
			panic(err)
		}
	}
	return nil
}
```

```go
// parseOne parses one flag. It reports whether a flag was seen.
func (f *FlagSet) parseOne() (bool, error) {
	if len(f.args) == 0 {
		return false, nil
	}
	s := f.args[0]
	if len(s) < 2 || s[0] != '-' {
		return false, nil
	}
	numMinuses := 1
	if s[1] == '-' {
		numMinuses++
		if len(s) == 2 { // "--" terminates the flags
			f.args = f.args[1:]
			return false, nil
		}
	}
	name := s[numMinuses:]
	if len(name) == 0 || name[0] == '-' || name[0] == '=' {
		return false, f.failf("bad flag syntax: %s", s)
	}

	// it's a flag. does it have an argument?
	f.args = f.args[1:]
	hasValue := false
	value := ""
	for i := 1; i < len(name); i++ { // equals cannot be first
		if name[i] == '=' {
			value = name[i+1:]
			hasValue = true
			name = name[0:i]
			break
		}
	}

	flag, ok := f.formal[name]
	if !ok {
		if name == "help" || name == "h" { // special case for nice help message.
			f.usage()
			return false, ErrHelp
		}
		return false, f.failf("flag provided but not defined: -%s", name)
	}

	if fv, ok := flag.Value.(boolFlag); ok && fv.IsBoolFlag() { // special case: doesn't need an arg
		if hasValue {
			if err := fv.Set(value); err != nil {
				return false, f.failf("invalid boolean value %q for -%s: %v", value, name, err)
			}
		} else {
			if err := fv.Set("true"); err != nil {
				return false, f.failf("invalid boolean flag %s: %v", name, err)
			}
		}
	} else {
		// It must have a value, which might be the next argument.
		if !hasValue && len(f.args) > 0 {
			// value is the next arg
			hasValue = true
			value, f.args = f.args[0], f.args[1:]
		}
		if !hasValue {
			return false, f.failf("flag needs an argument: -%s", name)
		}
		if err := flag.Value.Set(value); err != nil {
			return false, f.failf("invalid value %q for flag -%s: %v", value, name, err)
		}
	}
	if f.actual == nil {
		f.actual = make(map[string]*Flag)
	}
	f.actual[name] = flag
	return true, nil
}
```

即当控制台有参数时才更新数据

### 为什么默认值输出°C

```go
// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	// CommandLine 是默认的命令行标志集，从 os.Args 解析。顶级函数（如 [BoolVar]、[Arg] 等）是 CommandLine 方法的包装器。
	flag.CommandLine.Var(&f, name, usage)
	// flag.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")
```

`func CelsiusFlag(name string, value Celsius, usage string) *Celsius`
注意 value的类型是 Celsius，可以在类型安全的情况下，将默认值赋给 f 

这时候 如果控制台里没有标志位，`Set()` 就不会触发更新数据，此时就是用的 &f.Celsius 里面的数据

fmt.Println(*temp) 此时用的就是 *temp(伪代码：*(&f.Celsius)) 

→ 检查 *temp 类型为 Celsius
→ 发现实现了 String()
→ 调用 temp.String()



### 重新声明一个符合 flag.value接口 的类型 和 利用结构体构造符合

```go
// 生成符合flag.Value接口的time.Duration
// -- time.Duration Value
type durationValue time.Duration

func newDurationValue(val time.Duration, p *time.Duration) *durationValue {
	*p = val
	return (*durationValue)(p)
}

func (d *durationValue) Set(s string) error {
	v, err := time.ParseDuration(s)
	if err != nil {
		err = errParse
	}
	*d = durationValue(v)
	return err
}

func (d *durationValue) String() string { return (*time.Duration)(d).String() }

//...
```

```go
// *celsiusFlag satisfies the flag.Value interface.
// celsiusFlag内嵌了一个Celsius类型（§2.5），因此不用实现本身就已经有String方法了。
type celsiusFlag struct{ Celsius }
```

一个是重新声明了一个类型，基于新类型增加Set方法，以期符合flag.value

另一个是利用结构体包裹了初始类型，为其加上Set方法

这么看，*time.Duration 也可以用结构体构造，单独返回属性的地址



### 结构体占内存吗， 属性可以单独取地址？
// ch5/10-recover/01-title/title.go

https://chat.deepseek.com/a/chat/s/69fe564f-210d-4231-8520-323de956bcbb
为什么 用 type bailout struct{} 作为类型标志呢

bailout 是一个空结构体类型，它不携带任何数据（零内存开销）

那么 如果 
p:= bailout {}  , 这个有申请内存吗
零内存是什么意思？连地址都没有吗？

如果 我有个结构体是
x := struct {
    y int
}{
    y: 12
}
这个x结构体 占内存吗
还是只有属性单独占内存和内存地址，跟结构体没关系？