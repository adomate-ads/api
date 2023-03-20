package google_ads_controller

import (
	"encoding/json"
)

type Campaign struct {
	CustomerId     uint   `json:"customer_id,omitempty"`
	CampaignName   string `json:"campaign_name,omitempty"`
	CampaignBudget uint   `json:"campaign_budget,omitempty"`
	ResourceName   string `json:"resource_name,omitempty"`
}

func CreateCampaign(campaign Body) (*Campaign, error) {
	msg := Message{
		Route: "/create_campaign",
		Body:  campaign,
	}

	resp := SendToQueue(msg)
	var camp Campaign
	err := json.Unmarshal([]byte(resp), &camp)
	if err != nil {
		return nil, err
	}
	return &camp, nil
}

func GetCampaigns(campaign Body) ([]Campaign, error) {
	msg := Message{
		Route: "/get_campaigns",
		Body:  campaign,
	}

	resp := SendToQueue(msg)
	var campaigns []Campaign
	err := json.Unmarshal([]byte(resp), &campaigns)
	if err != nil {
		return nil, err
	}
	return campaigns, nil
}

func EnableCampaign(campaign Body) (*Campaign, error) {
	msg := Message{
		Route: "/enable_campaign",
		Body:  campaign,
	}

	resp := SendToQueue(msg)
	var camp Campaign
	err := json.Unmarshal([]byte(resp), &camp)
	if err != nil {
		return nil, err
	}
	return &camp, nil
}

func PauseCampaign(campaign Body) (*Campaign, error) {
	msg := Message{
		Route: "/pause_campaign",
		Body:  campaign,
	}

	resp := SendToQueue(msg)
	var camp Campaign
	err := json.Unmarshal([]byte(resp), &camp)
	if err != nil {
		return nil, err
	}
	return &camp, nil
}

func RemoveCampaign(campaign Body) (*Campaign, error) {
	msg := Message{
		Route: "/remove_campaign",
		Body:  campaign,
	}

	resp := SendToQueue(msg)
	var camp Campaign
	err := json.Unmarshal([]byte(resp), &camp)
	if err != nil {
		return nil, err
	}
	return &camp, nil
}
