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
