package google_ads_controller

import (
	"encoding/json"
)

type ReportResponse struct {
	Name             string  `json:"name"`
	Impressions      uint    `json:"impressions"`
	Clicks           uint    `json:"clicks"`
	ClickThroughRate float64 `json:"click_through_rate"`
	AverageCPC       float64 `json:"average_cpc"`
	CostMicros       uint    `json:"cost_micros"`
	Enabled          bool    `json:"enabled"`
}

func GetReport(customerId uint) (*[]ReportResponse, error) {
	msg := Message{
		Route: "/report",
		Body: Body{
			CustomerId: customerId,
		},
	}

	resp := SendToGAC(msg)
	var reports []ReportResponse
	err := json.Unmarshal([]byte(resp), &reports)
	if err != nil {
		return nil, err
	}
	return &reports, nil
}
