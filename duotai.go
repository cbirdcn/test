package main

import "fmt"

// 定义接口
type Animal interface {
    Walk()    // 函数原型
}

// 定义接收类型
type Human struct {
    name string
}
// 此接收类型实现接口要求的所有方法
func (h *Human) Walk() {
    fmt.Println(h.name)
}

type Dog struct {
    name string
}

func (d *Dog) Walk() {
    fmt.Println(d.name)
}

// 多态：传入不同的类对象，调用同样的方法，实现不同的效果
func DoWalk(a Animal) {
    a.Walk()
}

func main() {

	h := Human{"andy"}
	d := Dog{"d"}
	DoWalk(&h)
	DoWalk(&d)
}
