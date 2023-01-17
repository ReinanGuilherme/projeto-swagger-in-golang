package router

import (
	docs "SwaggerInGolang/docs" // Especificar o diretório raiz do projeto.

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	controller "SwaggerInGolang/src/controllers"
)

func SetupRoutes(r *gin.Engine) {

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")

	v1.GET("/HelloWorld", controller.HelloWorld)

	v1.GET("/HelloFriend", controller.HelloFriend)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
