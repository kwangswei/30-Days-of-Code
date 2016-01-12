package main

import "fmt"

type Person struct {
	age int
}

func (p Person) NewPerson(_age int) Person {
	if _age < 0 {
		fmt.Println("This person is not valid, setting age to 0.")
		p.age = 0
	} else {
		p.age = _age
	}
	return p
}

func (p Person) amIOld() {
	if p.age < 13 {
		fmt.Println("You are young.")
	} else if p.age < 18 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are old.")
	}
}

func (p Person) yearPasses() Person {
	p.age += 1
	return p
}

func main() {
	var T, age int

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&age)
		p := Person{age: age}
		p = p.NewPerson(age)
		p.amIOld()
		for j := 0; j < 3; j++ {
			p = p.yearPasses()
		}
		p.amIOld()
		fmt.Println()
	}
}
