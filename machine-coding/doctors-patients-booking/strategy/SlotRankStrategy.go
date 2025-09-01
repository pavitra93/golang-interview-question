package strategy

import "github.com/pavitra93/machine-coding/doctors-patients-bookings/DTO"

type SlotRankStrategy interface {
	SlotRank(slots []DTO.DoctorSlot)
}
