package google_ads_controller

import (
	"encoding/json"
)

type AdGroup struct {
	CustomerId   uint   `json:"customer_id,omitempty"`
	CampaignId   uint   `json:"campaign_id,omitempty"`
	AdGroupId    uint   `json:"ad_group_id,omitempty"`
	AdGroupName  string `json:"ad_group_name,omitempty"`
	ResourceName string `json:"resource_name,omitempty"`
	MinCPCBid    uint   `json:"min_cpc_bid,omitempty"`
}

func CreateAdGroup(adGroup Body) (*AdGroup, error) {
	msg := Message{
		Route: "/create_ad_group",
		Body:  adGroup,
	}

	resp := SendToQueue(msg)
	var adGrp AdGroup
	err := json.Unmarshal([]byte(resp), &adGrp)
	if err != nil {
		return nil, err
	}
	return &adGrp, nil
}

func GetAdGroups(adGroup Body) ([]AdGroup, error) {
	msg := Message{
		Route: "/get_ad_groups",
		Body:  adGroup,
	}

	resp := SendToQueue(msg)
	var adGroups []AdGroup
	err := json.Unmarshal([]byte(resp), &adGroups)
	if err != nil {
		return nil, err
	}
	return adGroups, nil
}

func EnableAdGroup(adGroup Body) (*AdGroup, error) {
	msg := Message{
		Route: "/enable_ad_group",
		Body:  adGroup,
	}

	resp := SendToQueue(msg)
	var adGrp AdGroup
	err := json.Unmarshal([]byte(resp), &adGrp)
	if err != nil {
		return nil, err
	}
	return &adGrp, nil
}

func PauseAdGroup(adGroup Body) (*AdGroup, error) {
	msg := Message{
		Route: "/pause_ad_group",
		Body:  adGroup,
	}

	resp := SendToQueue(msg)
	var adGrp AdGroup
	err := json.Unmarshal([]byte(resp), &adGrp)
	if err != nil {
		return nil, err
	}
	return &adGrp, nil
}

func RemoveAdGroup(adGroup Body) (*AdGroup, error) {
	msg := Message{
		Route: "/remove_ad_group",
		Body:  adGroup,
	}

	resp := SendToQueue(msg)
	var adGrp AdGroup
	err := json.Unmarshal([]byte(resp), &adGrp)
	if err != nil {
		return nil, err
	}
	return &adGrp, nil
}
