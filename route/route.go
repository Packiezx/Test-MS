package route

import (
	"templategoapi/controller"
	"templategoapi/db"

	"github.com/gin-gonic/gin"
)

//MemberRoute create route
func NewRoute(r *gin.Engine, resource *db.Resource) {

	api := r.Group("/api")
	{
		api.GET("/", controller.Test(resource))
		api.POST("/login", controller.Login(resource))
		api.POST("/accounts", controller.CreateAccount(resource))
	}
}
