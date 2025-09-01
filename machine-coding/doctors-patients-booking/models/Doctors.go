package models

import (
	"math/rand"

	"github.com/pavitra93/machine-coding/doctors-patients-bookings/enums/Specialization"
)

type Doctor struct {
	ID             int
	DoctorName     string
	Specialization Specialization.Specialization
	Availability   map[string]bool
	Rating         float64
}

func NewDoctor(name string, specialization Specialization.Specialization, rating float64) *Doctor {
	return &Doctor{
		ID:             rand.Intn(100),
		DoctorName:     name,
		Specialization: specialization,
		Rating:         rating,
	}
}
