package main

import "fmt"

// import "fmt"

// func increment() func() int {
// 	val := 0

// 	return func() int {
// 		val++
// 		return val
// 	}
// }

// func main() {
// 	a := increment()
// 	fmt.Println(a())
// 	fmt.Println(a())
// 	fmt.Println(a())
// }

func fact(n int) int{
	if n == 0{
		return 1
	}
	return n*fact(n-1)
}

func main(){
	myfact := fact(1)
	fmt.Println(myfact)
}