package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	IsRemote bool   `json:"isRemote"`
	Address  Address
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

func (e *Employee) updateName(newName string) {
	e.Name = newName
}

func (e *Employee) printAddress() {
	fmt.Println(e.Address)
}

func main() {
	address := Address{
		Street: "123 St.",
		City:   "New York",
	}
	employee1 := Employee{
		Name:     "Alice",
		Age:      30,
		IsRemote: true,
		Address:  address,
	}
	employee1.updateName("Bob")
	fmt.Println(employee1)
	employee1.printAddress()

	jsonData, _ := json.MarshalIndent(employee1, "", "   ")
	fmt.Println(string(jsonData))

	// anonymous struct
	job := struct {
		title  string
		salary int
	}{
		title:  "Software Engineer",
		salary: 250000,
	}
	fmt.Println(job)
}
