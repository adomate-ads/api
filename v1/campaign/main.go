package campaign

import (
	"github.com/adomate-ads/api/middleware/auth"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.POST("/campaign", auth.NotGuest, CreateCampaign)
	r.GET("/campaign", auth.NotGuest, GetCampaigns)
	r.GET("/campaign/company/:id", auth.NotGuest, GetCampaignsForCompany)
	r.GET("/campaign/:id", auth.NotGuest, GetCampaign)
	r.DELETE("/campaign/:id", auth.NotGuest, DeleteCampaign)
}
