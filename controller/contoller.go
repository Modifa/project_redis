package controller

import (
	"fmt"
	"net/http"

	models "github.com/Modifa/project_redis.git/models"
	services "github.com/Modifa/project_redis.git/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func UserTestRedis(c *gin.Context) {
	// db := s.DB{}
	//
	// var rb models.Returnblock
	var u models.User
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	//SaveTest
	err := services.SaveTest(u)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//Get All Donations

/* Get All Donations...*/
func GetAllDonations(c *gin.Context) {
	db := services.DB{}
	// var rb models.Returnblock
	var u models.EmptyRequest
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//use function
	response, _ := db.GetAllDonations("project_fin.get_all_donations", u)
	//LPush Transactions
	for i := 0; i < len(response); i++ {
		services.SaveTransactions(response[i])
	}
}
