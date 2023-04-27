package main

import (
	"control-habitos/internal/core/service"
	"control-habitos/internal/infra/primary/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	//Instancias
	serviceInstance := service.CreateAgregarHabitoServiceInstance()
	habitosController := controller.CreateHabitosController(serviceInstance)

	router := gin.Default()
	router.POST("/habitos/v1/crear", habitosController.Crear)
	router.Run(":8080")
}
