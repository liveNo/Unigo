package main

import "fmt"

// Go语言类型转换，类型转换用于将一种数据类型的变量转换为另外一种类型变量。
// Go语言类型转换基本格式如下： type_name(expression)

func main() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)

	fmt.Printf("mean 的值为：%f\n", mean)
}
