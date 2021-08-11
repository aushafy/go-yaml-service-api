package app

import (
	"github.com/aushafy/go-yaml-service-api/controllers"
)

func endpoints() {
	router.POST("/api/v1/trigger", controllers.Trigger)
}
