package campaign

import (
	"github.com/adomate-ads/api/middleware/auth"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.POST("/adgroup", auth.NotGuest, CreateCampaign)
	r.GET("/adgroup", auth.NotGuest, auth.InGroup("super-admin"), GetCampaigns)
	r.GET("/adgroup/company/:id", auth.NotGuest, GetCampaignsForCompany)
	r.GET("/adgroup/:id", auth.NotGuest, GetCampaign)
	r.DELETE("/adgroup/:id", auth.NotGuest, DeleteCampaign)
}
