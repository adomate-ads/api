package gads

import (
	"github.com/adomate-ads/api/middleware/auth"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	googleAds := r.Group("/gads")
	googleAds.GET("/client", auth.NotGuest, auth.InGroup("super-admin"), GetClients)
	googleAds.GET("/client/:clientId", auth.NotGuest, GetClient)

	googleAds.GET("/campaigns/", auth.NotGuest, auth.InGroup("super-admin"), GetCampaigns)
	googleAds.GET("/campaigns/:clientId", auth.NotGuest, GetCampaignsInClient)
	googleAds.GET("/campaign/:clientId/:campaignId", auth.NotGuest, GetCampaign)

	googleAds.GET("/adgroup/:clientId/:campaignId", auth.NotGuest, auth.InGroup("super-admin"), GetAdGroupsInCampaign)
	googleAds.GET("/adgroup/:clientId/:campaignId/:adgroupId", auth.NotGuest, GetAdGroup)

	googleAds.GET("/adgroupad/:clientId/:campaignId/:adgroupId", auth.NotGuest, GetAdGroupAds)

	googleAds.GET("/keyword/:clientId/:campaignId/:adgroupId", auth.NotGuest, GetKeywords)
}
