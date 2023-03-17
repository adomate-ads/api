package google_ads_controller

import (
	"encoding/json"
	"fmt"
)

type AdGroup struct {
	CustomerId  uint   `json:"customer_id,omitempty"`
	CampaignId  uint   `json:"campaign_id,omitempty"`
	AdGroupId   uint   `json:"ad_group_id,omitempty"`
	AdGroupName string `json:"ad_group_name,omitempty"`
	MinCPCBid   string `json:"min_cpc_bid,omitempty"`
}

func CreateAdGroup(adGroup AdGroup) {
	adGrp, err := json.Marshal(adGroup)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := Message{
		Route: "/create_ad_group",
		Body:  string(adGrp),
	}

	resp := SendToQueue(msg)
	fmt.Println(resp)
}

func GetAdGroups(adGroup AdGroup) {
	adGrp, err := json.Marshal(adGroup)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := Message{
		Route: "/get_ad_groups",
		Body:  string(adGrp),
	}

	resp := SendToQueue(msg)
	fmt.Println(resp)
}

func EnableAdGroup(adGroup AdGroup) {
	adGrp, err := json.Marshal(adGroup)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := Message{
		Route: "/enable_ad_group",
		Body:  string(adGrp),
	}

	resp := SendToQueue(msg)
	fmt.Println(resp)
}

func PauseAdGroup(adGroup AdGroup) {
	adGrp, err := json.Marshal(adGroup)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := Message{
		Route: "/pause_ad_group",
		Body:  string(adGrp),
	}

	resp := SendToQueue(msg)
	fmt.Println(resp)
}

func RemoveAdGroup(adGroup AdGroup) {
	adGrp, err := json.Marshal(adGroup)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := Message{
		Route: "/remove_ad_group",
		Body:  string(adGrp),
	}

	resp := SendToQueue(msg)
	fmt.Println(resp)
}
