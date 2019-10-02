// https://play.golang.org/p/VNh81g9R7FG
// Number:  1

package main

import "fmt"

type myStruct struct {
	number int
}

func main() {
	n := myStruct{}.number
	n += 1

	fmt.Println("Number: ", n)
}
