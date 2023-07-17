package google_ads_controller

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
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

func GetCampaignID(url string) (uint, error) {
	p := strings.TrimPrefix(url, "/")
	parts := strings.Split(p, "/")
	if len(parts) < 4 {
		return 0, fmt.Errorf("url does not match the expected structure")
	}
	id, err := strconv.ParseUint(parts[3], 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
