package DTO

import "github.com/pavitra93/machine-coding/doctors-patients-bookings/models"

type DoctorSlot struct {
	Doctor *models.Doctor
	Slot   string
}
