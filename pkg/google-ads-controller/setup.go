package google_ads_controller

import (
	"encoding/json"
	"fmt"
	"github.com/adomate-ads/api/pkg/discord"
	"github.com/adomate-ads/api/pkg/rabbitmq"
	"os"
)

type Message struct {
	Route string `json:"route" example:"/get_customers"`
	Body  Body   `json:"body,omitempty" example:"{'customer_name': 'Test Customer'}"`
}

type Body struct {
	Id           uint   `json:"id,omitempty"`
	CustomerName string `json:"customer_name,omitempty"`
	//Campaign
	CustomerId     uint   `json:"customer_id,omitempty"`
	CampaignName   string `json:"campaign_name,omitempty"`
	CampaignBudget uint   `json:"campaign_budget,omitempty"`
	//Ad Group Ads
	AdGroupId    uint     `json:"ad_group_id,omitempty"`
	Headlines    []string `json:"headlines,omitempty"`
	Descriptions []string `json:"descriptions,omitempty"`
	FinalURL     string   `json:"final_url,omitempty"`
	//Ad Group
	CampaignId  uint   `json:"campaign_id,omitempty"`
	AdGroupName string `json:"ad_group_name,omitempty"`
	MinCPCBid   uint   `json:"min_cpc_bid,omitempty"`
}

var Queue string = os.Getenv("RABBIT_GAC_QUEUE")

func SendToGAC(message Message) string {
	msgBody, err := json.Marshal(message)
	if err != nil {
		discord.SendMessage(discord.Error, "[GAC] Failed to marshal message", fmt.Sprintf("Error: %s", err.Error()))
		return ""
	}

	resp, err := rabbitmq.SendMessageWithResponse(msgBody, Queue)
	if err != nil {
		discord.SendMessage(discord.Error, "[GAC] Failed to send message", fmt.Sprintf("Error: %s", err.Error()))
		return ""
	}

	return resp
}
