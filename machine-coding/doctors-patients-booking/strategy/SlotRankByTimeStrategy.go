package strategy

import (
	"fmt"
	"sort"

	"github.com/pavitra93/machine-coding/doctors-patients-bookings/dto"
)

type SlotRankByTimeStrategy struct {
}

func (s *SlotRankByTimeStrategy) SlotRank(slots []dto.DoctorSlot) {
	sort.Slice(slots, func(i, j int) bool {
		return slots[j].Slot > slots[i].Slot
	})
	for i := 0; i < len(slots); i++ {
		fmt.Println(slots[i].Doctor.DoctorName, "==> ", slots[i].Slot)
	}
}
