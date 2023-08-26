package google_ads_controller

import (
	"encoding/json"
)

type Customer struct {
	Id            uint   `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Budget        string `json:"budget,omitempty"`
	Status        uint   `json:"status,omitempty"`
	ServingStatus uint   `json:"serving_status,omitempty"`
	ResourceName  string `json:"resource_name,omitempty"`
	Enabled       bool   `json:"enabled,omitempty"`
}

func CreateCustomer(customerName string) (*Customer, error) {
	msg := Message{
		Route: "/create_customer",
		Body: Body{
			CustomerName: customerName,
		},
	}

	resp := SendToGAC(msg)
	var customer Customer
	err := json.Unmarshal([]byte(resp), &customer)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func GetCustomers() ([]Customer, error) {
	msg := Message{
		Route: "/get_customers",
	}

	resp := SendToGAC(msg)
	var customers []Customer
	err := json.Unmarshal([]byte(resp), &customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}
