// go_study_1
package main

import (
	"fmt"
)

func main() {
	fmt.Println("**** test array ****")
	arrayTest()
}
func arrayTest() {
	var arr [10]int
	arr[0] = 10
	arr[1] = 20
	fmt.Printf("the first element in arr is :%d\n", arr[0])
}
