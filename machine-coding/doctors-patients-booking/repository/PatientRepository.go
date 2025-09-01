package repository

import (
	"github.com/pavitra93/machine-coding/doctors-patients-bookings/models"
)

type PatientRepository struct {
	PatientMap map[int]*models.Patients
}

func NewPatientRepository() *PatientRepository {
	return &PatientRepository{
		PatientMap: make(map[int]*models.Patients),
	}
}

func (repo *PatientRepository) Save(patient *models.Patients) (*models.Patients, error) {
	if repo.PatientMap == nil {
		repo.PatientMap = make(map[int]*models.Patients)
	}

	repo.PatientMap[patient.ID] = patient

	return patient, nil
}

func (repo *PatientRepository) FindByID(ID int) (*models.Patients, bool) {

	patient, exists := repo.PatientMap[ID]

	return patient, exists
}

func (repo *PatientRepository) FindAll() []*models.Patients {
	patients := make([]*models.Patients, 0, len(repo.PatientMap))
	for _, patient := range repo.PatientMap {
		patients = append(patients, patient)
	}
	return patients
}
