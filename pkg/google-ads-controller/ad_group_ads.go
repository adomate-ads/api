package google_ads_controller

import (
	"encoding/json"
	"fmt"
)

type AdGroupAds struct {
	CustomerId   string   `json:"customer_id,omitempty"`
	AdGroupId    string   `json:"ad_group_id,omitempty"`
	Headlines    []string `json:"headlines,omitempty"`
	Descriptions []string `json:"descriptions,omitempty"`
	FinalURL     []string `json:"final_url,omitempty"`
}

func CreateAdGroupAds(adGroupAds AdGroupAds) {
	adGrpAds, err := json.Marshal(adGroupAds)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := Message{
		Route: "/create_ad",
		Body:  string(adGrpAds),
	}

	resp := SendToQueue(msg)
	fmt.Println(resp)
}

func GetAdGroupAds(adGroupAds AdGroupAds) {
	adGrpAds, err := json.Marshal(adGroupAds)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := Message{
		Route: "/get_ads",
		Body:  string(adGrpAds),
	}

	resp := SendToQueue(msg)
	fmt.Println(resp)
}

func RemoveAdGroupAd(adGroupAds AdGroupAds) {
	adGrpAds, err := json.Marshal(adGroupAds)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := Message{
		Route: "/remove_ad",
		Body:  string(adGrpAds),
	}

	resp := SendToQueue(msg)
	fmt.Println(resp)
}
