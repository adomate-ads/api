package google_ads_controller

import (
	"encoding/json"
)

type AdGroupAds struct {
	CustomerId   uint     `json:"customer_id,omitempty"`
	AdGroupId    uint     `json:"ad_group_id,omitempty"`
	Headlines    []string `json:"headlines,omitempty"`
	Descriptions []string `json:"descriptions,omitempty"`
	FinalURL     string   `json:"final_url,omitempty"`
	ResourceName string   `json:"resource_name,omitempty"`
}

func CreateAdGroupAds(adGroupAds Body) (*AdGroupAds, error) {
	msg := Message{
		Route: "/create_ad",
		Body:  adGroupAds,
	}

	resp := SendToQueue(msg)
	var adGrpAd AdGroupAds
	err := json.Unmarshal([]byte(resp), &adGrpAd)
	if err != nil {
		return nil, err
	}
	return &adGrpAd, nil
}

func GetAdGroupAds(adGroupAds Body) ([]AdGroupAds, error) {
	msg := Message{
		Route: "/get_ads",
		Body:  adGroupAds,
	}

	resp := SendToQueue(msg)
	var adGrpAds []AdGroupAds
	err := json.Unmarshal([]byte(resp), &adGrpAds)
	if err != nil {
		return nil, err
	}
	return adGrpAds, nil
}

func RemoveAdGroupAd(adGroupAds Body) (*AdGroupAds, error) {
	msg := Message{
		Route: "/remove_ad",
		Body:  adGroupAds,
	}

	resp := SendToQueue(msg)
	var adGrpAd AdGroupAds
	err := json.Unmarshal([]byte(resp), &adGrpAd)
	if err != nil {
		return nil, err
	}
	return &adGrpAd, nil
}
