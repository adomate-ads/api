package user

import (
	"github.com/adomate-ads/api/middleware/auth"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.POST("/user", auth.NotGuest, auth.InGroup("admin"), CreateUser)
	r.GET("/user", auth.NotGuest, auth.InGroup("super-admin"), GetUsers)
	r.GET("/user/company/:id", auth.NotGuest, auth.InGroup("admin"), GetUsersByCompanyID)
	r.GET("/user/:id", auth.NotGuest, auth.InGroup("admin"), GetUser)
	r.PATCH("/user/:user", auth.NotGuest, UpdateUser)
	r.DELETE("/user/:id", auth.NotGuest, auth.InGroup("admin"), DeleteUser)
}
