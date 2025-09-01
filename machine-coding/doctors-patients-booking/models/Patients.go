package models

import "math/rand"

type Patients struct {
	ID   int
	Name string
}

func NewPatient(name string) *Patients {
	return &Patients{rand.Intn(100), name}
}
