package main

// Employee结构体
import (
	"time"
)

type Employee struct {
	ID           int
	Name         string
	Address      string
	Age          int
	Sex          bool
	Hiredate     time.Time
	Job          int
	ManagerId    int
	DepartmentId int
	test         int
}
type Department struct {
	ID       int
	Name     string
	Location string
}
