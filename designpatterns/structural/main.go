package main

import (
	"fmt"
	flyweight "github.com/pavitra93/golang-interview-questions.git/designpatterns/structural/flyweight"
)

func main() {
	fmt.Println("Welcome to Flyweight DP")

	bulletBuilder := flyweight.
		GetInstance().
		SetRadius(6.0).
		SetLength(9.0).
		SetBulletType(flyweight.Hollow_Point).
		SetMetal(flyweight.IRON).
		Build()

	flyingBullet1 := flyweight.FlyingBullet{X: 6.0, Y: 7.0, Z: 9.0, Bullet: bulletBuilder}
	flyingBullet2 := flyweight.FlyingBullet{X: 6.0, Y: 7.0, Z: 9.0, Bullet: bulletBuilder}
	flyingBullet3 := flyweight.FlyingBullet{X: 6.0, Y: 7.0, Z: 9.0, Bullet: bulletBuilder}

	fmt.Println(flyingBullet1)
	fmt.Println(flyingBullet2)
	fmt.Println(flyingBullet3)

}
