package main

import "fmt"

func main() {
	score := 89
	switch {
	case score >= 90:
		fmt.Printf("优秀")
	case score >= 80:
		fmt.Printf("良好")
	case score >= 60:
		fmt.Printf("及格")
	default:
		fmt.Printf("差")
	}
}
