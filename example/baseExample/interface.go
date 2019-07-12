package main

import "fmt"

// Go语言接口： Go语言提供了另外一种数据类型即接口。它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you")
}

type Iphone struct{}

func (iphone Iphone) call() {
	fmt.Println("I am Iphone, I can call you")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(Iphone)
	phone.call()
}
