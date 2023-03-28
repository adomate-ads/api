package helpers

import (
	"github.com/adomate-ads/api/models"
	"github.com/adomate-ads/api/pkg/discord"
	"strconv"
)

func SyncClient(clientId string) {
	// Get all campaigns inside client
	campaigns := GetCampaigns(clientId)
	// Compare campaigns to database
	for _, campaign := range campaigns {
		c, err := models.GetCampaignByGoogleID(uint(campaign.Id))
		if err != nil {
			// if campaign is not in database, create it
			companyIdInt, err := strconv.ParseInt(clientId, 10, 64)
			if err != nil {
				discord.SendMessage("error", "Error running sync.", "NA")
				return
			}
			company, err := models.GetCompanyByClientID(companyIdInt)
			if err != nil {
				discord.SendMessage("error", "Error running sync.", "NA")
				return
			}

			c := models.Campaign{
				ResourceName: campaign.ResourceName,
				GoogleID:     uint(campaign.Id),
				CompanyID:    company.ID,
				Company:      *company,
			}

			err = c.CreateCampaign()
			if err != nil {
				discord.SendMessage("error", "Error running sync.", "NA")
				return
			}
		} else {
			// if campaign is in database, update it
			SyncCampaign(clientId, strconv.Itoa(int(c.GoogleID)))
		}
	}
	// TODO - This doesnt check for rogue campaigns in the database
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
				discord.SendMessage("error", "Error running sync.", "NA")
				return
			}
			campaign, err := models.GetCampaignByGoogleID(uint(campaignIdInt))
			if err != nil {
				discord.SendMessage("error", "Error running sync.", "NA")
				return
			}

			companyIdInt, err := strconv.ParseInt(clientId, 10, 64)
			if err != nil {
				discord.SendMessage("error", "Error running sync.", "NA")
				return
			}
			company, err := models.GetCompanyByClientID(companyIdInt)
			if err != nil {
				discord.SendMessage("error", "Error running sync.", "NA")
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
				discord.SendMessage("error", "Error running sync.", "NA")
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
