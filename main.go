package main

import (
	"fmt"
	solid "github.com/pavitra93/golang-interview-questions.git/solid"
)

func main() {
	fmt.Println("Welcome to SOLID principles")

	Eagle := &solid.Eagle{}
	Eagle.SetName("Eagle")
	Eagle.SetColor("red")
	Eagle.SetSize("L")
	Eagle.SetBirdType(solid.EAGLE)
	Eagle.Fly()
	fmt.Println(Eagle.GetSize())

	Parrot := &solid.Parrot{}
	Parrot.SetName("Parrot")
	Parrot.SetColor("red")
	Parrot.SetSize("L")
	Parrot.SetBirdType(solid.PARROT)
	Parrot.Fly()
	fmt.Println(Parrot.GetBirdType())

	Peguin := &solid.Peguin{}
	Peguin.SetName("Peguin")
	Peguin.SetColor("red")
	Peguin.SetSize("L")
	Peguin.SetBirdType(solid.PEGUINE)
	Peguin.Swim()
	Peguin.SetNonCommonAttribute("Non CommonAttribute")
	fmt.Println(Peguin.GetNonCommonAttribute())

	Duck := &solid.Duck{}
	Duck.SetName("Duck")
	Duck.SetColor("red")
	Duck.SetSize("L")
	Duck.SetBirdType(solid.DUCK)
	Duck.Fly()
	Duck.Swim()
	fmt.Println(Duck.GetBirdType())
}
