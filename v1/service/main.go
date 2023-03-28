package service

import (
	"github.com/adomate-ads/api/middleware/auth"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.POST("/service", auth.NotGuest, auth.InGroup("super-admin"), CreateService)
	r.GET("/service", auth.NotGuest, auth.InGroup("super-admin"), GetServices)
	r.GET("/service/company/:id", auth.NotGuest, auth.InGroup("admin"), GetServicesForCompany)
	r.GET("/service/:id", auth.NotGuest, GetService)
	r.PATCH("/service/:id", auth.NotGuest, auth.InGroup("super-admin"), UpdateService)
	r.DELETE("/service/:id", auth.NotGuest, auth.InGroup("super-admin"), DeleteService)
}
