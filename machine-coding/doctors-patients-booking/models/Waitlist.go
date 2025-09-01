package models

// Waitlist model maintains a FIFO queue of patient IDs per (doctorID, slot string).
type Waitlist struct {
	Queues map[string][]int // key: doctorID|slotString -> value: queue of patientIDs
}

func NewWaitlist() *Waitlist {
	return &Waitlist{
		Queues: make(map[string][]int),
	}
}
