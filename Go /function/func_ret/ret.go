package main

import "time"

type Employee struct {
	ID      int
	Name    string
	Address string
	DoB     time.Time
}

func EmployeeByID(id int) Employee {
	return Employee{ID: id}
}
func main() {
	// EmployeeByID(1).Name = "H"
	// error:
	// EmployeeById(int)返回的是值类型，值类型不能被赋值，只有变量能被赋值。
	var a = EmployeeByID(1)
	a.Name = "H"
}
