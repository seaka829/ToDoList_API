package router

import (
	"../controller"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := router()
	r.Run()

}

func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/api/v1/tasks")
	{
		ctrl := controller.Controller{}
		u.GET("", ctrl.ReadAll)
		u.GET("/:id", ctrl.ReadByID)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.UpdateByID)
		u.DELETE("/:id", ctrl.DeleteByID)
	}

	return r
}
