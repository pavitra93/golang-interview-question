package strategy

import "github.com/pavitra93/machine-coding/doctors-patients-bookings/dto"

type SlotRankStrategy interface {
	SlotRank(slots []dto.DoctorSlot)
}
