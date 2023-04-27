package controller

import (
	"control-habitos/internal/core/domain"
	"control-habitos/internal/core/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

var (
	once sync.Once
	habitosController *HabitosController
)

func CreateHabitosController(agregarHabito usecase.AgregarHabito) *HabitosController {
	once.Do(func() {
		habitosController = &HabitosController{agregarHabito: agregarHabito}
	})
	return habitosController
}

type HabitosController struct {
	agregarHabito usecase.AgregarHabito
}

func (h *HabitosController) Crear(c *gin.Context)  {
	var request domain.Habito
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error en el request"})
		return
	}

	if request.Titulo == "" || request.Horario == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Titulo u Horario vacios. Reintentelo"})
		return
	}

	habito, err := h.agregarHabito.Handle(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Un error inesperado a ocurrido. Intentelo nuevamente"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensaje": "Habito creado con exito",
		"ID": habito.Id,
	})
}


