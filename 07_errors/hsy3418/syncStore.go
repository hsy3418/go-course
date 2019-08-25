package main

import (
	"fmt"
	"sync"
)

type SyncStore struct {
	sync.Mutex
	sync.Map
}

func NewSyncStore() *SyncStore {
	return &SyncStore{}
}

// CreatePuppy adds a nuw puppy to the puppies store
func (m *SyncStore) CreatePuppy(puppy Puppy) error {
	m.Lock()
	defer m.Unlock()
	if puppy.Value < 0 {
		err := &Error{
			Message: "The puppy value is invalidate",
			Code:    -1,
		}
		return err
	}
	if _, exists := m.Load(puppy.ID); exists {
		return fmt.Errorf("puppy with %d ID already exists", puppy.ID)
	}
	m.Store(puppy.ID, puppy)
	return nil
}

// ReadPuppy retrieves the puppy for a given id from puppies store
func (m *SyncStore) ReadPuppy(id int32) (Puppy, error) {
	if p, exists := m.Load(id); exists {
		puppy, _ := p.(Puppy)
		return puppy, nil
	}
	return Puppy{}, &Error{Message: "puppy ID does not exist", Code: -2}
}

//UpdatePuppy updates the puppy for the given id
func (m *SyncStore) UpdatePuppy(puppy Puppy) error {
	m.Lock()
	defer m.Unlock()
	if puppy.Value < 0 {
		err := &Error{
			Message: "The puppy value is invalidate",
			Code:    -1,
		}
		return err
	}
	if _, exists := m.Load(puppy.ID); !exists {
		return &Error{Message: "puppy ID does not exist", Code: -2}
	}
	m.Store(puppy.ID, puppy)
	return nil
}

//DeletePuppy delete the puppy for the given id from puppies store
func (m *SyncStore) DeletePuppy(id int32) error {
	m.Lock()
	defer m.Unlock()
	if _, exists := m.Load(id); exists {
		m.Delete(id)
		return nil
	}
	return &Error{Message: "puppy ID does not exist", Code: -2}
}
