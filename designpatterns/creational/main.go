package main

import (
	"fmt"
	builder "github.com/pavitra93/golang-interview-questions.git/designpatterns/creational/builder"
	factory "github.com/pavitra93/golang-interview-questions.git/designpatterns/creational/factory"
	prototype "github.com/pavitra93/golang-interview-questions.git/designpatterns/creational/prototype"
	singleton "github.com/pavitra93/golang-interview-questions.git/designpatterns/creational/singleton"
)

func main() {
	fmt.Println("Welcome to Singleton Design Pattern")
	singleObject := singleton.GetInstance()
	fmt.Println(singleObject)

	singleObject2 := singleton.GetInstance()
	fmt.Printf("%T\n", singleObject2)

	fmt.Println(singleObject == singleObject2)

	// ============================
	fmt.Println("Welcome to Builder Design Pattern")
	ParameterizedConstructorObject := builder.NewStudent("Pavitra", "pavitramehta4@gmail.com", "8233564291")
	fmt.Println(ParameterizedConstructorObject) // Immuatable

	NonParameterizedConstructorObject := &builder.Student{}
	NonParameterizedConstructorObject.SetName("Pavitra")
	NonParameterizedConstructorObject.SetEmail("pavitramehta4@gmail.com")
	fmt.Println(NonParameterizedConstructorObject)

	// Builder Pattern
	User1Object := builder.
		GetInstance().
		SetName("Pavitra").
		SetEmail("pavitramehta4@gmail.com").
		SetPhone("8233564291").
		SetActive().
		Build()
	fmt.Println(User1Object)

	User2Object := builder.
		GetInstance().
		SetName("Pavitra").
		SetEmail("pavitramehta4@gmail.com").
		SetActive().
		Build()
	fmt.Println(User2Object)

	// ============================
	fmt.Println("Welcome to Prototype Design Pattern")

	register := prototype.NewRegistry()

	Original := &prototype.Rectangle{Red: 1.0, Green: 2.0, Blue: 3.0}
	register.Register("rectangle", Original)
	fmt.Println("Original:", Original)

	shape, _ := register.Get("rectangle")
	clone1 := shape.(*prototype.Rectangle)
	clone1.Blue = 8.0
	register.Register("rectangle2", clone1)
	fmt.Println("Clone 1:", clone1)

	clone2 := shape.(*prototype.Rectangle)
	clone2.Red = 11.45
	clone2.Green = 12.0
	clone2.Blue = 13.0
	register.Register("rectangle3", clone2)
	fmt.Println("Clone 2:", clone2)

	for k, v := range register.RegisterMap {
		fmt.Println(k)
		fmt.Println(v)
	}

	// ==============FACTORY DESIGN PATTERN==============
	ButtonFactory := &factory.ButtonFactory{}
	PrimaryButton := ButtonFactory.CreateButton(factory.PRIMARY)
	PrimaryButton.ClickButton()

	SecondaryButton := ButtonFactory.CreateButton(factory.SECONDARY)
	SecondaryButton.ClickButton()

}
