package delivery

import (
	docs "otus/internal/delivery/swagger/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title otus homeWork
// @version 1.0
// @description otus homeWork
// @license.name kanya384

// @contact.name API Support
// @contact.email kanya384@mail.ru

// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func (d *Delivery) initRouter() *gin.Engine {

	var router = gin.New()

	d.routerDocs(router.Group("/docs"))
	d.routerUsers(router.Group("/user"))

	return router
}

func (d *Delivery) routerUsers(router *gin.RouterGroup) {
	router.GET("/:id", d.ReadUserById)
	router.PUT("/", d.CreateUser)
	router.POST("/", d.UpdateUser)
	router.DELETE("/:id", d.DeleteUser)
}

func (d *Delivery) routerDocs(router *gin.RouterGroup) {
	docs.SwaggerInfo.BasePath = "/"

	router.Any("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
