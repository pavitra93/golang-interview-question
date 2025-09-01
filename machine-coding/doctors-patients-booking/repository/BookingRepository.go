package repository

import "github.com/pavitra93/machine-coding/doctors-patients-bookings/models"

type BookingRepository struct {
	BookingMap map[int]*models.Bookings
}

func NewBookingRepository() *BookingRepository {
	return &BookingRepository{
		BookingMap: make(map[int]*models.Bookings),
	}
}

func (repo *BookingRepository) Save(booking *models.Bookings) (*models.Bookings, error) {
	repo.BookingMap[booking.BookingID] = booking

	return booking, nil
}

func (repo *BookingRepository) FindByID(ID int) (*models.Bookings, bool) {
	booking, exists := repo.BookingMap[ID]
	return booking, exists
}

func (repo *BookingRepository) FindAll() []*models.Bookings {
	bookings := make([]*models.Bookings, 0, len(repo.BookingMap))
	for _, booking := range repo.BookingMap {
		bookings = append(bookings, booking)
	}
	return bookings
}

func (repo *BookingRepository) Delete(ID int) {
	delete(repo.BookingMap, ID)
	return
}

func (repo *BookingRepository) FindByDoctor(ID int) []*models.Bookings {
	bookings := make([]*models.Bookings, 0)
	for _, booking := range repo.BookingMap {
		if booking.DoctorID == ID {
			bookings = append(bookings, booking)
		}
	}
	return bookings
}

func (repo *BookingRepository) FindByPatient(ID int) []*models.Bookings {
	bookings := make([]*models.Bookings, 0)
	for _, booking := range repo.BookingMap {
		if booking.PatientID == ID {
			bookings = append(bookings, booking)
		}
	}
	return bookings
}
