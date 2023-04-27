package service

import (
	"control-habitos/internal/core/domain"
	"fmt"
	"sync"
)

var (
	once sync.Once
	agregarHabitoService *AgregarHabitoService
)

type AgregarHabitoService struct {}

func CreateAgregarHabitoServiceInstance() *AgregarHabitoService {
	once.Do(func() {
		agregarHabitoService = &AgregarHabitoService{}
	})
	return agregarHabitoService
}

func (a *AgregarHabitoService) Handle(habito domain.Habito) (*domain.Habito, error) {
	fmt.Println("Agregando un habito")

	return &domain.Habito{
		Id:      "",
		Titulo:  "Funcionando",
		Horario: "",
	}, nil
}
