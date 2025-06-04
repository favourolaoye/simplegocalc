package main

import (
	"fmt"
	"github.com/favourolaoye/simplegocalc/calculator"
)
var x float32
var y float32
var operation string

var score float32
var StudID string

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
	fmt.Println("Enter the studentID: ")
	fmt.Scanln(&StudID)
	fmt.Println("Enter the Grade: ")
	fmt.Scanln(&score)

	grade:= calculator.Grader(score, StudID)

	fmt.Printf("Student with ID: %v grade is: %v  and their data has been saved!", StudID, grade)
}