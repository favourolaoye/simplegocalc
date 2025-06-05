package main

import (
	"fmt"
)

var num int;

func iterations(elt int, arr []int) bool{
	for _, v := range arr{
		if v == elt{
			return true
		}
	}
	return false
}

func main(){
	fmt.Println("Enter the num to search for: ")
	fmt.Scanln(&num)
	element := iterations(num, []int{1,2,3,4,5,6,7})
	fmt.Printf("status ->  %v", element)	
} 