package singleton

import (
	"fmt"
	"sync"
)

type database struct{}

var instance *database
var once sync.Once

func GetInstance() *database {
	once.Do(func() {
		fmt.Println("Database instance not initialized")
		instance = &database{}
	})

	fmt.Println("Database instance initialized")
	return instance
}
