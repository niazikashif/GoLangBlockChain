package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

type Student struct {
	rollnumber int
	name       string
	address    string
	subject    []string
}

func CalculateHash(stringToHash string) string {
	fmt.Printf("String Received : %s\n", stringToHash)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
}

func NewStudent(rollno int, name string, address string, subject []string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	s.subject = subject
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string, subject []string) *Student {
	st := NewStudent(rollno, name, address, subject)
	ls.list = append(ls.list, st)
	return st
}

func (ls *StudentList) DisplayList() {
	fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), 1, strings.Repeat("=", 25))
	fmt.Printf("student rollno %d\nstudent name %s\nstudent address %s\nstudent subjects %v\n", ls.list[0].rollnumber, ls.list[0].name, ls.list[0].address, ls.list[0].subject[:])
	fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), 2, strings.Repeat("=", 25))
	fmt.Printf("student rollno %d\nstudent name %s\nstudent address %s\nstudent subjects %v\n", ls.list[1].rollnumber, ls.list[1].name, ls.list[1].address, ls.list[1].subject[:])

}

func main() {
	student := new(StudentList)
	student.CreateStudent(24, "asim", "AAAAAA", []string{"Blockchain", "Info Security"})
	student.CreateStudent(25, "Naveed", "BBBBBB", []string{"OOP", "AI"})
	student.DisplayList()

	fmt.Println("\n\n")

	for _, val := range student.list {
		fmt.Printf("Hash: %x\n", CalculateHash(val.name))
	}

}
