package main

import "fmt"

func main() {
	var name string
	fmt.Scanln(&name)
	var gender string
	fmt.Scanln(&gender)
	var age int
	fmt.Scanln(&age)

	var newIntern = intern{
		name:   name,
		gender: gender,
		age:    age,
	}

	fmt.Println(newIntern)
}

type intern struct {
	name   string
	gender string
	age    int
}
