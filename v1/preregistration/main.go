package preregistration

import (
	"github.com/gin-gonic/gin"
)

// TODO Log IPs so we can log which IP created what domain and only allow them to edit it
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
