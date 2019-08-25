package main

import "fmt"

// MapStore is a implementation of storer for the storage of puppies
type MapStore struct {
	puppies map[int32]Puppy
}

func NewMapStore() *MapStore {
	return &MapStore{
		puppies: map[int32]Puppy{},
	}
}

// CreatePuppy adds a nuw puppy to the puppies store
func (m *MapStore) CreatePuppy(puppy Puppy) error {
	if puppy.Value < 0 {
		err := &Error{
			Message: "The puppy value is invalidate",
			Code:    -1,
		}
		return err
	}
	if _, exists := m.puppies[puppy.ID]; !exists {
		m.puppies[puppy.ID] = puppy
		return nil
	}
	return fmt.Errorf("puppy with %d ID already exists", puppy.ID)
}

// ReadPuppy retrieves the puppy for a given id from puppies store
func (m *MapStore) ReadPuppy(id int32) (Puppy, error) {
	if _, exists := m.puppies[id]; !exists {
		return Puppy{}, &Error{Message: "puppy ID does not exist", Code: -2}
	}
	return m.puppies[id], nil
}

//UpdatePuppy updates the puppy for the given id
func (m *MapStore) UpdatePuppy(puppy Puppy) error {
	if puppy.Value < 0 {
		err := &Error{
			Message: "The puppy value is invalidate",
			Code:    -1,
		}
		return err
	}
	if _, exists := m.puppies[puppy.ID]; !exists {
		return &Error{Message: "puppy ID does not exist", Code: -2}
	}
	m.puppies[puppy.ID] = puppy
	return nil
}

//DeletePuppy delete the puppy for the given id from puppies store
func (m *MapStore) DeletePuppy(id int32) error {
	if _, exists := m.puppies[id]; exists {
		delete(m.puppies, id)
		return nil
	}
	return &Error{Message: "puppy ID does not exist", Code: -2}
}
