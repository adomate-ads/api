package preregistration

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.POST("/preregistration", CreatePreRegistration)
	r.POST("/preregistration/locations", AddLocations)
	r.POST("/preregistration/services", AddServices)
	r.POST("/preregistration/budget", SetBudget)
	r.DELETE("/preregistration/locations", DeleteLocations)
	r.DELETE("/preregistration/services", DeleteServices)
}
