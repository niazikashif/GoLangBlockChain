package main

import "fmt"

type Str struct {
	Name string
}

func foo(var_ Str) {
	fmt.Println(var_.Name)
}
func main() {

	foo(Str{Name: "Niazi"})
}
