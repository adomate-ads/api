package helpers

import (
	"github.com/adomate-ads/api/models"
	"strconv"
)

func SyncClient(clientId string) {
}

func SyncCampaign(clientId, campaignId string) {
	// Get all adgroups inside campaign
	adGroups := GetAdGroups(clientId, campaignId)
	// Compare adgroups to database
	for _, adGroup := range adGroups {
		ag, err := models.GetAdGroupByGoogleID(uint(adGroup.Id))
		if err != nil {
			// if adgroup is not in database, create it
			campaignIdInt, err := strconv.ParseInt(campaignId, 10, 32)
			if err != nil {
				// TODO - Internal Server Error Panic to discord
				return
			}
			campaign, err := models.GetCampaignByGoogleID(uint(campaignIdInt))
			if err != nil {
				// TODO - Internal Server Error Panic to discord
				return
			}

			companyIdInt, err := strconv.ParseInt(clientId, 10, 64)
			if err != nil {
				// TODO - Internal Server Error Panic to discord
				return
			}
			company, err := models.GetCompanyByClientID(companyIdInt)
			if err != nil {
				// TODO - Internal Server Error Panic to discord
				return
			}

			ag := models.AdGroup{
				Name:         adGroup.Name,
				ResourceName: adGroup.ResourceName,
				GoogleID:     uint(adGroup.Id),
				CampaignID:   campaign.ID,
				Campaign:     *campaign,
				CompanyID:    company.ID,
				Company:      *company,
			}

			err = ag.CreateAdGroup()
			if err != nil {
				// TODO - Internal Server Error Panic to discord
				return
			}
		} else {
			// if adgroup is in database, update it
			SyncAdGroup(ag.GoogleID)
		}
	}
	// TODO - This doesnt check for rogue adgroups in the database
}

func SyncAdGroup(adGroupId uint) {

}
