package usecase

import "control-habitos/internal/core/domain"

type AgregarHabito interface {
	Handle(habito domain.Habito) (*domain.Habito, error)
}
