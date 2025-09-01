package service

import (
	"fmt"

	"github.com/pavitra93/machine-coding/doctors-patients-bookings/DTO"
	"github.com/pavitra93/machine-coding/doctors-patients-bookings/enums/Specialization"
	"github.com/pavitra93/machine-coding/doctors-patients-bookings/models"
	"github.com/pavitra93/machine-coding/doctors-patients-bookings/repository"
	"github.com/pavitra93/machine-coding/doctors-patients-bookings/strategy"
)

type BookingService struct {
	BookingRepository *repository.BookingRepository
	PatientRepository *repository.PatientRepository
	DoctorRepository  *repository.DoctorRepository
	WaitingRepository *repository.WaitlistRepository
}

func (service *BookingService) ViewBookingsByDoctor(ID int) ([]*models.Bookings, error) {
	return service.BookingRepository.FindByDoctor(ID), nil
}

func (service *BookingService) ViewBookingsByPatient(ID int) ([]*models.Bookings, error) {
	return service.BookingRepository.FindByPatient(ID), nil
}

func (service *BookingService) Book(patientId int, doctorId int, slot string) (int, error) {

	_, exists := service.PatientRepository.FindByID(patientId)
	if !exists {
		return 0, fmt.Errorf("Patient with ID %d not found\n", patientId)
	}

	doctor, exists := service.DoctorRepository.FindByID(doctorId)
	if !exists {
		return 0, fmt.Errorf("Doctor with ID %d not found\n", doctorId)
	}

	if service.WaitingRepository.CheckPatientAlreadyInQueueForDoctorID(patientId, doctorId, slot) {
		return 0, fmt.Errorf("patient with ID %d is already in queue for doctor with ID %d", patientId, doctorId)
	}

	bookings := service.BookingRepository.FindByPatient(patientId)
	for _, booking := range bookings {
		if booking.Slot == slot {
			return 0, fmt.Errorf("patient with ID %d is already in booked for slot %s", patientId, slot)
		}
	}

	avail, exists := service.DoctorRepository.CheckAvailabilitySlotAvailable(doctorId, slot)
	if exists {
		if avail {
			booking := models.NewBooking(patientId, doctorId, slot)
			service.BookingRepository.Save(booking)
			doctor.Availability[slot] = false
			return booking.BookingID, nil
		} else {
			service.WaitingRepository.Add(doctorId, slot, patientId)
			fmt.Printf("Patient with ID %d is added to waitlist for doctor with ID %d\n", patientId, doctorId)
		}
	} else {
		fmt.Printf("Doctor with ID %d is not available for slot %s\n", doctorId, slot)
	}

	return 0, nil

}

func (service *BookingService) Cancel(bookingID int) {
	booking, exists := service.BookingRepository.FindByID(bookingID)
	if !exists {
		fmt.Printf("Booking with ID %d not found\n", bookingID)
	}

	service.BookingRepository.Delete(bookingID)
	doctor, exists := service.DoctorRepository.FindByID(booking.DoctorID)
	if !exists {
		fmt.Printf("Doctor with ID %d not found\n", bookingID)
	}
	doctor.Availability[booking.Slot] = true
	fmt.Printf("Booking with ID %d is deleted successfully\n", bookingID)

	waitListPatientId, exists := service.WaitingRepository.Pop(booking.DoctorID, booking.Slot)
	if !exists {
		fmt.Printf("No Patient in queue for slot\n", booking.Slot)
	}

	NewbookingID, _ := service.Book(waitListPatientId, booking.DoctorID, booking.Slot)
	fmt.Printf("Booking with ID %d is created successfully\n", NewbookingID)

}

func (service *BookingService) List(specialization Specialization.Specialization, strategy strategy.SlotRankStrategy) {
	doctors := service.DoctorRepository.FindBySpecialization(specialization)

	var results []DTO.DoctorSlot

	for _, doctor := range doctors {
		for slot, avail := range doctor.Availability {
			if avail {
				results = append(results, DTO.DoctorSlot{
					Doctor: doctor,
					Slot:   slot,
				})
			}
		}
	}

	strategy.SlotRank(results)
}
