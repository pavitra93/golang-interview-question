package main

import (
	"fmt"

	"github.com/pavitra93/machine-coding/doctors-patients-bookings/enums/Specialization"
	"github.com/pavitra93/machine-coding/doctors-patients-bookings/repository"
	"github.com/pavitra93/machine-coding/doctors-patients-bookings/service"
	"github.com/pavitra93/machine-coding/doctors-patients-bookings/strategy"
)

func main() {
	fmt.Println("Welcome to Doctors & Patients Booking System")

	DoctorRepo := repository.NewDoctorRepository()
	PatientRepo := repository.NewPatientRepository()

	DoctorService := service.DoctorService{DoctorRepository: DoctorRepo}
	PatientService := service.PatientService{PatientRepository: PatientRepo}

	// Doctors Registration
	doctorPavitra, _ := DoctorService.Register("Dr. Pavitra Mehta", Specialization.CARDIOLOGIST, 4.5)
	doctorPavitra, _ = DoctorService.DeclareAvailability(doctorPavitra.ID, []string{"10:00", "11:00", "12:00"})
	fmt.Printf("%#v\n", doctorPavitra)

	doctorMonika, _ := DoctorService.Register("Dr. Monika Mehta", Specialization.CARDIOLOGIST, 4.0)
	doctorMonika, _ = DoctorService.DeclareAvailability(doctorMonika.ID, []string{"14:00", "15:00", "16:00"})
	fmt.Printf("%#v\n", doctorMonika)

	// Patients Registration
	Patient1, _ := PatientService.Register("Nitesh Mehta")
	Patient2, _ := PatientService.Register("Sudha Mehta")
	Patient3, _ := PatientService.Register("Dilip Mehta")

	WaitRepo := repository.NewWaitlistRepository()
	BookingRepo := repository.NewBookingRepository()
	BookingService := service.BookingService{
		BookingRepository: BookingRepo,
		DoctorRepository:  DoctorRepo,
		PatientRepository: PatientRepo,
		WaitingRepository: WaitRepo,
	}

	booking1, err := BookingService.Book(Patient1.ID, doctorPavitra.ID, "10:00")
	if err != nil {
		fmt.Println(err)
	}
	_, err = BookingService.Book(Patient2.ID, doctorPavitra.ID, "11:00")
	if err != nil {
		fmt.Println(err)
	}

	_, err = BookingService.Book(Patient3.ID, doctorPavitra.ID, "12:00")
	if err != nil {
		fmt.Println(err)
	}

	_, err = BookingService.Book(Patient3.ID, doctorMonika.ID, "10:00")
	if err != nil {
		fmt.Println(err)
	}

	_, err = BookingService.Book(Patient3.ID, doctorMonika.ID, "15:00")
	if err != nil {
		fmt.Println(err)
	}

	waitlist := WaitRepo.List()
	for k, v := range waitlist {
		fmt.Println(k, v)
	}

	BookingService.Cancel(booking1)

	waitlist = WaitRepo.List()
	for k, v := range waitlist {
		fmt.Println(k, v)
	}

	bookings, _ := BookingService.ViewBookingsByDoctor(doctorPavitra.ID)
	for k, v := range bookings {
		fmt.Println(k, v)
	}

	bookings, _ = BookingService.ViewBookingsByDoctor(doctorMonika.ID)
	for k, v := range bookings {
		fmt.Println(k, v)
	}

	ratingStrategy := &strategy.SlotRankByRatingStrategy{}
	BookingService.List(Specialization.CARDIOLOGIST, ratingStrategy)

	fmt.Println("=============================")

	timeStrategy := &strategy.SlotRankByTimeStrategy{}
	BookingService.List(Specialization.CARDIOLOGIST, timeStrategy)

}
