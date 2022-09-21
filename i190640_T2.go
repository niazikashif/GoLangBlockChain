package main

import "fmt"

type company struct {
	companyName string
	employees   []employee
}

type employee struct {
	name     string
	salary   int
	position string
}

func main() {
	emp1 := employee{"Amir", 80000, "Full-Stack Developer"}
	emp2 := employee{"Niazi", 90000, "Data Scientist"}
	emp3 := employee{"Ahmed", 10000, "Block Chain Developer"}

	emplys := []employee{emp1, emp2, emp3}
	company := company{"Tetra", emplys}

	fmt.Println("Company Details : ", company)
}
