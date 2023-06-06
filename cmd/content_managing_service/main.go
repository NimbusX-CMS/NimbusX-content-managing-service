package main

import (
	"fmt"
	"github.com/NimbusX-CMS/NimbusX-content-managing-service/internal/api"
	"github.com/NimbusX-CMS/NimbusX-content-managing-service/internal/bussniss_logic"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	server := &bussniss_logic.Server{}

	api.RegisterHandlers(router, server)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
