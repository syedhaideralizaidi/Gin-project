// main.go

package main

import (
	"github.com/username/go-gin-postgresql-backend/src/models"
	"github.com/username/go-gin-postgresql-backend/src/routes"
	"github.com/username/go-gin-postgresql-backend/src/utils"
)

func main() {
	utils.LoadEnv()
	models.OpenDatabaseConnection()
	models.AutoMigrateModels()
	router := routes.SetupRoutes()
	router.Run(":8080")

}
