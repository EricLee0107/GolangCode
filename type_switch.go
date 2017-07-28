package main

import "fmt"

func main() {
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型是：%T", i)
	case int:
		fmt.Printf("x 是int型")
	case float64:
		fmt.Printf("x是float64型")
	case string:
		fmt.Printf("x是string型")
	default:
		fmt.Printf("未知型")
	}
}
