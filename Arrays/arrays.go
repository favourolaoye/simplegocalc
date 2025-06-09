// package main

// import "fmt"

// func arrays(items[]int, val int) []int{
// 	// num := [5]int{1, 2, 3, 4, 5}
// 	// fmt.Println(num)
// 	joined := append(items, val)
// 	return joined

// }

// func main(){
// 	created := arrays([]int {1,2,4,5}, 12)
// 	fmt.Println(created);

// 	test := [5]int{1,2,3,4,5}

// 	for _, v := range test{
// 		fmt.Println(v)
// 	} 
// 	// this use the key, value --> it unwraps the index and elet directly

// 	// for i := 0; i < len(test); i++{
// 	// 	fmt.Println(test[i])
// 	// }
// 	// this is an old way, but not bad also
// }


// slices
package main

import "fmt"

func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0

    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

    sum(1, 2)
    sum(1, 2, 3)

    nums := []int{1, 2, 3, 4}
    sum(nums...)
}