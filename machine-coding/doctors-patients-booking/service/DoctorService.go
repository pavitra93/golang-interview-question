package service

import (
	"fmt"

	"github.com/pavitra93/machine-coding/doctors-patients-bookings/enums/Specialization"
	"github.com/pavitra93/machine-coding/doctors-patients-bookings/models"
	"github.com/pavitra93/machine-coding/doctors-patients-bookings/repository"
)

type DoctorService struct {
	DoctorRepository *repository.DoctorRepository
}

func (service *DoctorService) Register(name string, specialization Specialization.Specialization, rating float64) (*models.Doctor, error) {
	doctor := models.NewDoctor(name, specialization, rating)
	service.DoctorRepository.Save(*doctor)
	return doctor, nil
}

func (service *DoctorService) DeclareAvailability(ID int, slots []string) (*models.Doctor, error) {
	doctor, ok := service.DoctorRepository.FindByID(ID)
	if !ok {
		fmt.Printf("Doctor with ID %d not found\n", ID)
	}

	if doctor.Availability == nil {
		doctor.Availability = make(map[string]bool)
	}

	for _, slot := range slots {
		doctor.Availability[slot] = true
	}

	return doctor, nil
}
