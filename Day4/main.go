package main

import "fmt"

type person struct {
	age int
}

func NewPerson(initialAge int) person {
	p := person{age: initialAge}
	//Add some more code to run some checks on initialAge
	p.age = initialAge
	if p.age < 0 {
		p.age = 0
		fmt.Println("Age is not valid, setting age to 0.")
	}
	return p
}

func (p *person) amIOld() {
	//Do some computation in here and print out the correct statement to the console
	switch {
	case p.age < 13:
		fmt.Println("You are young.")
	case p.age >= 13 && p.age < 18:
		fmt.Println("You are a teenager.")
	default:
		fmt.Println("You are old.")
	}
}

func (p *person) yearPasses() {
	//Increment the age of the person in here
	p.age++
}

func main() {
	var T, age int

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&age)
		p := NewPerson(age)
		p.amIOld()
		for j := 0; j < 3; j++ {
			p.yearPasses()
		}
		p.amIOld()
		fmt.Println()
	}
}
