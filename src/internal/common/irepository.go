package common

import "github.com/google/uuid"

type IRepository interface {
	Get(id uuid.UUID) (interface{}, error)
	Create(entity interface{}) error
	Update(id uuid.UUID, entity interface{}) error
	Delete(id uuid.UUID) error
}
