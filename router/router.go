package router

import (
	"fast-project-golang/controller"
	"fast-project-golang/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db",db)
	})
	public     :=  r.Group("/api")
	public.POST("/authentication", controller.MiddlewareAuth)
	public.POST("/register",controller.RegisterAuth)

	protected  := r.Group("/api")
	protected.Use(middleware.JwtAuthMiddleware())
	protected.GET("/transaction",controller.FindTransaction)
	protected.POST("/transaction/create",controller.CreateTransaction)
	protected.PATCH("/transaction/:id",controller.UpdateTransaction)
	protected.DELETE("/transaction/:id",controller.DeleteTransaction)

	return r
}