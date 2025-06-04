package main

import (
	"fmt"
	"golang/calculator"
)
var x float32
var y float32
var operation string

func main() {
	fmt.Println("Hey welcome to webflux calculator (lol)!")
	fmt.Println("-------------------------------------------------------------------------------------")
	fmt.Println("Enter the first number(whole number): ");
	fmt.Scanln(&x)
	fmt.Println("Enter the second number(whole number): ");
	fmt.Scanln(&y);
	fmt.Println("Enter the arithemetic operation you wanna perform(sum, diff, div, mul): ");
	fmt.Scanln(&operation)
	result:= calculator.Calc(x, y, operation)
	fmt.Printf("The %v is: %v", operation , result)
}