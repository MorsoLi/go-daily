package main

import (
	"flag"
	"fmt"
)

var (
	intflag    *int
	boolflag   *bool
	stringflag *string
)

func init() {
	// 调用flag.Type（其中Type可以为Int/Uint/Bool/Float64/String/Duration等）
	// 会自动为我们分配变量，返回该变量的地址
	boolflag = flag.Bool("bool", true, "bool value")
	intflag = flag.Int("intflag", 1, "int flag value")
	stringflag = flag.String("string", "string", "string value")
	// 使用flag.TypeVar定义选项，先定义变量，然后变量的地址
	//flag.IntVar(&intflag, "intflag", 0, "int flag value")
	//flag.BoolVar(&boolflag, "bool", false, "bool value")
	//flag.StringVar(&stringflag, "string", "default", "string value")
}

// support long args
func test_init() {
	var logVar string
	const defLogVal = "hahaa"
	const usage = "test long args"
	flag.StringVar(&logVar, "long_args", defLogVal, usage)
	flag.StringVar(&logVar, "l", defLogVal, usage+"shorthand")
}

func main() {
	flag.Parse()
	fmt.Println("int", *intflag)
	fmt.Println("str", *stringflag)
	fmt.Println("bool", *boolflag)
}
