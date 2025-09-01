package repository

import (
	"github.com/pavitra93/machine-coding/doctors-patients-bookings/enums/Specialization"
	"github.com/pavitra93/machine-coding/doctors-patients-bookings/models"
)

type DoctorRepository struct {
	doctorsMap map[int]*models.Doctor
}

func NewDoctorRepository() *DoctorRepository {
	return &DoctorRepository{
		doctorsMap: make(map[int]*models.Doctor),
	}
}

func (repo *DoctorRepository) Save(doctor models.Doctor) *models.Doctor {
	repo.doctorsMap[doctor.ID] = &doctor

	return &doctor

}

func (repo *DoctorRepository) FindByID(ID int) (*models.Doctor, bool) {
	doctor, ok := repo.doctorsMap[ID]

	return doctor, ok
}

func (repo *DoctorRepository) CheckAvailabilitySlotAvailable(doctorID int, slot string) (bool, bool) {
	avail, exists := repo.doctorsMap[doctorID].Availability[slot]

	return avail, exists
}

func (repo *DoctorRepository) FindAll() []*models.Doctor {
	doctors := make([]*models.Doctor, 0, len(repo.doctorsMap))
	for _, d := range repo.doctorsMap {
		doctors = append(doctors, d)
	}
	return doctors
}

func (repo *DoctorRepository) FindBySpecialization(specialization Specialization.Specialization) []*models.Doctor {
	doctors := make([]*models.Doctor, 0)
	for _, doctor := range repo.doctorsMap {
		if doctor.Specialization == specialization {
			doctors = append(doctors, doctor)
		}
	}

	return doctors
}
