package session

import (
	"training-service/training"
)

type Manager struct {
}

type Repository interface {
}

func (manager Manager) HasActiveSession(userId string) bool {
	return false
}

func (manager Manager) CreateSession(userId string, title string, location string) (training.Session, error) {
	return training.Session{}, nil
}

func (manager Manager) EndSession(userId string) error {
	return nil
}
