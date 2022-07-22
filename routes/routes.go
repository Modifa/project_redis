package routes

import (
	cont "github.com/Modifa/project_redis.git/controller"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	//reddis Synce Test Endpoint
	E1 := r.Group("/api/redis/test/")
	{
		E1.POST("User", cont.UserTestRedis)
	}
	E2 := r.Group("/api/Donations/")
	{
		E2.GET("GetAllDonations", cont.GetAllDonations)
	}
}
