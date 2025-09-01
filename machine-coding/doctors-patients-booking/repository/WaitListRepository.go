package repository

import (
	"fmt"
	"strings"
	"sync"

	"github.com/pavitra93/machine-coding/doctors-patients-bookings/models"
)

// WaitlistRepository maintains a FIFO queue of patient IDs per (doctorID, slot string).
type WaitlistRepository struct {
	mu       *sync.Mutex
	WaitList *models.Waitlist
}

func NewWaitlistRepository() *WaitlistRepository {
	return &WaitlistRepository{
		mu:       &sync.Mutex{},
		WaitList: models.NewWaitlist(),
	}
}

// WaitKey builds the key from doctorID and slot string (trimmed).
func WaitKey(doctorID int, slot string) string {
	return fmt.Sprintf("%d|%s", doctorID, strings.TrimSpace(slot))
}

func (r *WaitlistRepository) CheckPatientAlreadyInQueueForDoctorID(patientID int, doctorID int, slot string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := WaitKey(doctorID, slot)
	queue, ok := r.WaitList.Queues[key]
	if !ok || len(queue) == 0 {
		return false
	}
	for _, patient := range queue {
		if patient == patientID {
			return true
		}
	}

	return false
}

// Add adds a patientID to the tail of the waitlist for the given doctorID and slot.
func (r *WaitlistRepository) Add(doctorID int, slot string, patientID int) {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := WaitKey(doctorID, slot)
	r.WaitList.Queues[key] = append(r.WaitList.Queues[key], patientID)
}

// Pop removes and returns the patientID at the head of the waitlist for the given doctorID and slot.
// Returns (0, false) if the waitlist is empty or does not exist.
func (r *WaitlistRepository) Pop(doctorID int, slot string) (int, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := WaitKey(doctorID, slot)
	queue, ok := r.WaitList.Queues[key]
	if !ok || len(queue) == 0 {
		return 0, false
	}

	patientID := queue[0]
	queue = queue[1:]
	if len(queue) == 0 {
		delete(r.WaitList.Queues, key)
	} else {
		r.WaitList.Queues[key] = queue
	}
	return patientID, true
}

// Peek returns the patientID at the head without removing it.
func (r *WaitlistRepository) Peek(doctorID int, slot string) (int, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := WaitKey(doctorID, slot)
	queue, ok := r.WaitList.Queues[key]
	if !ok || len(queue) == 0 {
		return 0, false
	}
	return queue[0], true
}

// Len returns the current length of the waitlist for the given doctorID and slot.
func (r *WaitlistRepository) Len(doctorID int, slot string) int {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := WaitKey(doctorID, slot)
	return len(r.WaitList.Queues[key])
}

func (r *WaitlistRepository) List() []*models.Waitlist {
	return []*models.Waitlist{r.WaitList}
}
