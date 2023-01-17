package main

import (
	"github.com/gin-gonic/gin"

	routes "SwaggerInGolang/src/router"
)

// @title Gin Swagger Example API
// @version 1.0
// @description Este é um servidor de amostra.

// @contact.name Reinan Guilherme
// @contact.url https://www.linkedin.com/in/reinan-guilherme-34086b236

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
