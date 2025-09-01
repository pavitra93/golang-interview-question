package service

import (
	"github.com/pavitra93/machine-coding/doctors-patients-bookings/models"
	"github.com/pavitra93/machine-coding/doctors-patients-bookings/repository"
)

type PatientService struct {
	PatientRepository *repository.PatientRepository
}

func (service *PatientService) Register(name string) (*models.Patients, error) {
	patient := models.NewPatient(name)
	patient, _ = service.PatientRepository.Save(patient)
	return patient, nil
}

func (service *PatientService) FindByID(ID int) (*models.Patients, error) {
	patient, exists := service.PatientRepository.FindByID(ID)
	if !exists {
		return &models.Patients{}, nil
	}
	return patient, nil
}
