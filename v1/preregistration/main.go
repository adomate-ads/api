package preregistration

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.POST("/preregistration", CreatePreRegistration)
	r.GET("/preregistration", GetPreRegistrations)
	r.GET("/preregistration/:domain", GetPreRegistration)

	r.POST("/preregistration/locations", AddLocations)
	r.GET("/preregistration/locations", GetLocations)
	r.DELETE("/preregistration/locations", DeleteLocations)

	r.POST("/preregistration/services", AddServices)
	r.GET("/preregistration/services", GetServices)
	r.DELETE("/preregistration/services", DeleteServices)

	r.POST("/preregistration/budget", SetBudget)
}
