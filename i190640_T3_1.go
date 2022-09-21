package main

import (
	"fmt"
	"strings"
)

type Student struct {
	rollnumber int
	name       string
	address    string
}

func NewStudent(rollno int, name string, address string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string) *Student {
	st := NewStudent(rollno, name, address)
	ls.list = append(ls.list, st)
	return st
}

func (ls *StudentList) DisplayList() {
	fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), 1, strings.Repeat("=", 25))
	fmt.Printf("student rollno %d\nstudent name %s\nstudent address %s\n", ls.list[0].rollnumber, ls.list[0].name, ls.list[0].address)
	fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), 2, strings.Repeat("=", 25))
	fmt.Printf("student rollno %d\nstudent name %s\nstudent address %s\n", ls.list[1].rollnumber, ls.list[1].name, ls.list[1].address)

}

func main() {
	student := new(StudentList)
	student.CreateStudent(24, "asim", "AAAAAA")
	student.CreateStudent(25, "Naveed", "BBBBBB")
	student.DisplayList()

}
