package repositories

import (
	"context"
	"saude-mais/internal/domain/entities"
)

type PatientRepository interface {
	Create(ctx context.Context, patient *entities.Patient) error
	GetByEmail(ctx context.Context, email string) (*entities.Patient, error)
	GetByID(ctx context.Context, id int) (*entities.Patient, error)
	EmailExists(ctx context.Context, email string) (bool, error)
	Delete(id string) error
}
