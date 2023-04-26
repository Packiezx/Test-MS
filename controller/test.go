package controller

import (
	"templategoapi/db"

	"github.com/gin-gonic/gin"
)

func Test(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"SERVICE": "G2EPAY",
		})
	}
}
