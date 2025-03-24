package routes

import (
	controllers2 "easy-gin/app/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", controllers2.IndexHome)
	router.GET("/index", controllers2.IndexHome)

	router.GET("/users/:id", controllers2.UserGet)
	router.GET("/users", controllers2.UserGetList)
	router.POST("/users", controllers2.UserPost)
	router.PUT("/users/:id", controllers2.UserPut)
	router.PATCH("/users/:id", controllers2.UserPut)
	router.DELETE("/users/:id", controllers2.UserDelete)
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello Gin!"})
	})
	return r
}
