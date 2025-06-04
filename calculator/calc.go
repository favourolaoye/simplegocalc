package calculator

import "fmt"



func Calc(a float32, b float32, op string) float32 {
	if op == "sum"{
		return a + b
	}else if op == "diff"{
		return a - b
	}else if op == "div" {
		if b <= 0{
			return 0
		}
		return a / b
	}else if op == "mul" {
		return a * b
	}else {
		fmt.Println("Invalid operation")
		return 0
	}
}