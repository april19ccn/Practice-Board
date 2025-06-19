package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func EmployeeByID(id int) *Employee {
	dilbert := Employee{ID: id, Name: "Dilbert", Position: "Pointy-haired boss", Salary: 1000000, ManagerID: 0}
	return &dilbert
}

func main() {
	fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"

	id := dilbert.ID
	// EmployeeByID 改为值类型会报错：不能赋值给 EmployeeByID(id).Salary（既不能寻址，也不是映射索引表达式）
	EmployeeByID(id).Salary = 0 // fired for... no real reason
}
