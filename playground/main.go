package main

import "fmt"

type person struct {
	age int
}

func main() {
	people := make([]person, 2, 10)
	p1 := &people[1] // p1 := &person{}
	fmt.Println(p1)
	fmt.Printf("%p \n", p1)
	fmt.Printf("%p \n", &people[1])
	p1.age++
	fmt.Println(p1.age)
	fmt.Println(people[1].age)
	fmt.Println("len, cap", len(people), cap(people))
	//
	//// Add a new person
	people = append(people, person{})
	fmt.Println("len, cap", len(people), cap(people))
	//
	//// Add another year for p1
	p1.age++
	fmt.Printf("%p \n", p1)
	fmt.Printf("%p \n", &people[1])
	fmt.Println(people[1].age)
}
