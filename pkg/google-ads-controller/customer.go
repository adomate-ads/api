package google_ads_controller

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	Id           uint   `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Status       uint   `json:"status,omitempty"`
	ResourceName string `json:"resource_name,omitempty"`
}

func CreateCustomer(customerName string) {
	msg := Message{
		Route: "/create_customer",
		Body:  fmt.Sprintf("{'customer_name': '%s'}", customerName),
	}

	resp := SendToQueue(msg)
	fmt.Println(resp)
}

func GetCustomers() {
	msg := Message{
		Route: "/get_customers",
		Body:  "",
	}

	resp := SendToQueue(msg)
	var customers []Customer
	err := json.Unmarshal([]byte(resp), &customers)
	if err != nil {
		fmt.Println(err)
		return
	}
}
