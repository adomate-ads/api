package gads

import (
	"github.com/adomate-ads/api/middleware/auth"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	googleAds := r.Group("/gads")
	googleAds.GET("/client", auth.NotGuest, auth.InGroup("super-admin"), GetClients)
	googleAds.GET("/client/:clientId", auth.NotGuest, GetClient)

	googleAds.GET("/campaigns/:clientId", auth.NotGuest, auth.InGroup("super-admin"), GetCampaigns)
	googleAds.GET("/campaign/:clientId/:campaignId", auth.NotGuest, GetCampaign)
	//
	//googleAds.GET("/adgroup", auth.NotGuest, auth.InGroup("super-admin"), GetAdGroups)
	//googleAds.GET("/adgroup/:id", auth.NotGuest, GetAdGroup)
	//
	//googleAds.GET("/ad", auth.NotGuest, auth.InGroup("super-admin"), GetAds)
	//googleAds.GET("/ad/:id", auth.NotGuest, GetAd)
	//
}
