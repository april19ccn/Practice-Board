// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

// far := 212 包级不能用

var far = Fahrenheit(212)

// far = Fahrenheit(312) // 包级别不允许出现可执行语句

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func Compare() {
	// far = AbsoluteZeroC
	// cannot use AbsoluteZeroC (constant -273.15 of float64 type Celsius) as Fahrenheit value in assignment

	far = CToF(AbsoluteZeroC)
	fmt.Println(far)
}

func Conversion() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F

	// fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch
	// invalid operation: boilingF - FreezingC (mismatched types Fahrenheit and Celsius)
}

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
