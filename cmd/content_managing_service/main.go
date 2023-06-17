package main

import (
	"fmt"
	"github.com/NimbusX-CMS/NimbusX-content-managing-service/internal/api"
	"github.com/NimbusX-CMS/NimbusX-content-managing-service/internal/auth"
	"github.com/NimbusX-CMS/NimbusX-content-managing-service/internal/business_logic"
	"github.com/NimbusX-CMS/NimbusX-content-managing-service/internal/db/multi_db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db, err := multi_db.ConnectToSQLite("db.sqlite")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	err = db.EnsureTablesCreation()
	if err != nil {
		fmt.Println("Error creating tables:", err)
		return
	}
	server := &business_logic.Server{
		DB:   db,
		Auth: auth.Auth{DB: db},
	}

	api.RegisterHandlers(router, server)

	err = router.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
