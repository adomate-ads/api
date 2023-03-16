package google_ads_controller

import "fmt"

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
	fmt.Println(resp)
}
