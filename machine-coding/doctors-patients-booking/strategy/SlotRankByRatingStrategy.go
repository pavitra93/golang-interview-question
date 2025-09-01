package strategy

import (
	"fmt"
	"sort"

	"github.com/pavitra93/machine-coding/doctors-patients-bookings/DTO"
)

type SlotRankByRatingStrategy struct {
}

func (s *SlotRankByRatingStrategy) SlotRank(slots []DTO.DoctorSlot) {
	sort.Slice(slots, func(i, j int) bool {
		return slots[j].Doctor.Rating < slots[i].Doctor.Rating
	})

	for i := 0; i < len(slots); i++ {
		fmt.Println(slots[i].Doctor.DoctorName, "==> ", slots[i].Slot)
	}
}
