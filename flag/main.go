package main

import (
	"flag"
	"fmt"
)

// go run main.go -f -v test1 test2
// go run main.go -f test1 test2
func case1() {
	var showProgress, force bool
	// -f 当存在时拷贝，是否强制拷贝

	flag.BoolVar(&force, "f", false, "我是绑定的bool变量1")
	flag.BoolVar(&showProgress, "v", false, "我是绑定的bool变量2")
	flag.Parse()
	fmt.Println(flag.NArg())
	// 获取参数个数，必须要输入两个参数，因为copy是从这个文件到另一个文件
	if flag.NArg() < 2 {
		flag.Usage() // 打印用途
		return
	}
	fmt.Println(force, showProgress)
	fmt.Println(flag.Arg(0), flag.Arg(1))
}

func main() {
	case1()
}
