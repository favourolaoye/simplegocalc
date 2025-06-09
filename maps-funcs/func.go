// functions in golang

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func myfunc(param type, param type) return_type {
// code block
// must return something
// }

func randomizer(elements ...int) int{

	if len(elements) == 0 {
		return -1 
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano())) 
	return elements[r.Intn(len(elements))]

}
func main() {
	randomValue := randomizer(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Randomly selected value:", randomValue)
}