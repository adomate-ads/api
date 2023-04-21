package get_started

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.POST("/get-started", CreateAccount)
	r.GET("/get-started/location-services/:domain", GetLocationsAndServices)
}
