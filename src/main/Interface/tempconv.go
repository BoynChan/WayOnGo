package main

import (
	"flag"
	"fmt"
)

type Celsius float64 //摄氏温度

func (c Celsius) String() string {
	return fmt.Sprintf("%.1f°C", float64(c))
}

type Kelvins float64    //开尔文温度
type Fahrenheit float64 //华氏温度

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) String() string {
	return ""
}

//摄氏度转为华氏度
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

//华氏度转为摄氏度
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

//开尔文温度转为摄氏度
func KToC(k Kelvins) Celsius {
	return Celsius(k - 273.15)
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	//扫描字符串s,按照格式分为value和unit
	_, _ = fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K":
		f.Celsius = KToC(Kelvins(value))
		return nil
	}
	return fmt.Errorf("invalid temp %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
