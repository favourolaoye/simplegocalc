// functions in golang

package main

import "fmt"

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func myfunc(param type, param type) return_type {
// code block
// must return something
// }

// func randomizer(elements ...int) int{

// 	if len(elements) == 0 {
// 		return -1
// 	}

// 	rgen := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	return elements[rgen.Intn(len(elements))]

// }
// func main() {
// 	randomValue := randomizer(1, 2, 3, 4, 5, 6, 7, 8, 9)
// 	fmt.Println("Randomly selected value:", randomValue)
// }

func main(){
    myobject := map[string]int {
		"key": 20,
		"key2": 30,
	}
	fmt.Println(myobject)
	for key, value := range myobject{
		fmt.Printf("%v -> %v \n", key, value)
	}
}