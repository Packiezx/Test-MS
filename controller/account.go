package controller

import (
	"templategoapi/db"
	"templategoapi/model"
	"templategoapi/repo"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	return func(c *gin.Context) {
		var body Body
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"erro": err.Error(),
			})
			return
		}

		// test login
		var accountData model.AccountModel
		filter := bson.M{
			"username": body.Username,
			"password": body.Password,
			"count": bson.M{
				"$ne": 5,
			},
		}
		filterOption := bson.M{
			"datetime": -1,
		}

		err := repo.GetOneStatement(resource, "accounts", filter, filterOption, &accountData)
		if err != nil {
			c.JSON(400, gin.H{
				"code": 400,
			})
			return
		}

		c.JSON(200, gin.H{
			"code":    200,
			"payload": accountData,
		})
	}
}

func CreateAccount(resource *db.Resource) func(c *gin.Context) {
	type Body struct {
		Username string `bson:"username" json:"username"`
		Password string `bson:"password" json:"password"`
		Name     string `bosn:"name" json:"name"`
		Email    string `json:"email"`
	}
	return func(c *gin.Context) {
		var body Body
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"erro": err.Error(),
			})
			return
		}

		accountData := model.AccountModel{
			Datetime: time.Now(),
			Username: body.Username,
			Password: body.Password,
			Name:     body.Name,
		}

		result, err := repo.CreateStatement(resource, "accounts", accountData)
		if err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"erro": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"code":    200,
			"payload": result,
		})
	}
}
