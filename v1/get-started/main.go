package get_started

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.POST("/get-started", CreateAccount)
	//r.GET("/get-started/location-service/:domain", cloudflare.Verify, GetLocationsAndServices)
	r.POST("/get-started/location-service/:domain", GetLocationsAndServices)
	r.POST("/get-started/ad-content", GetAdContent)
	r.GET("/get-started/ip-info", GetIpInfo)
}
