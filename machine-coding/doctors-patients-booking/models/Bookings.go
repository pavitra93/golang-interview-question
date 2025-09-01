package models

import "math/rand"

type Bookings struct {
	BookingID int
	PatientID int
	DoctorID  int
	Slot      string
}

func NewBooking(patientId int, doctorId int, slot string) *Bookings {
	return &Bookings{rand.Int(), patientId, doctorId, slot}
}
