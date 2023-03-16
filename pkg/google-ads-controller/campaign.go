package google_ads_controller

import (
	"encoding/json"
	"fmt"
)

type Campaign struct {
	CustomerId     string `json:"customer_id,omitempty"`
	CampaignName   string `json:"campaign_name,omitempty"`
	CampaignBudget string `json:"campaign_budget,omitempty"`
}

func CreateCampaign(campaign Campaign) {
	camp, err := json.Marshal(campaign)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := Message{
		Route: "/create_campaign",
		Body:  string(camp),
	}

	resp := SendToQueue(msg)
	fmt.Println(resp)
}

func GetCampaigns(campaign Campaign) {
	camp, err := json.Marshal(campaign)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := Message{
		Route: "/get_campaigns",
		Body:  string(camp),
	}

	resp := SendToQueue(msg)
	fmt.Println(resp)
}

func EnableCampaign(campaign Campaign) {
	camp, err := json.Marshal(campaign)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := Message{
		Route: "/enable_campaign",
		Body:  string(camp),
	}

	resp := SendToQueue(msg)
	fmt.Println(resp)
}

func PauseCampaign(campaign Campaign) {
	camp, err := json.Marshal(campaign)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := Message{
		Route: "/pause_campaign",
		Body:  string(camp),
	}

	resp := SendToQueue(msg)
	fmt.Println(resp)
}

func RemoveCampaign(campaign Campaign) {
	camp, err := json.Marshal(campaign)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := Message{
		Route: "/remove_campaign",
		Body:  string(camp),
	}

	resp := SendToQueue(msg)
	fmt.Println(resp)
}
