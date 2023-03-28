package email

import (
	"html/template"
	"net/http"
)

type invoiceData struct {
	InvoiceNumber int
	CompanyName   string
}

type invoiceCosting struct {
	Product         string
	ProductPrice    float32
	InvoiceSubtotal float32
	InvoiceTax      float32
	InvoiceTotal    float32
}

type paymentDetails struct {
	CCNumber                   int
	BillingDetailsName         string
	BillingDetailsAddressLine1 string
	BillingDetailsAddressLine2 string
	BillingDetailsAddressState string
	BillingDetailsAddressZip   int
}

type emailData struct {
	InvoiceData    invoiceData
	InvoiceCosting invoiceCosting
	PaymentDetails paymentDetails
}

func invoiceEmailPageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("new-invoice.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	invoice := invoiceData{
		InvoiceNumber: 1934,
		CompanyName:   "Adomate LLC",
	}
	//TODO add support for multiple items
	//use {{.range}} in HTML to loop
	costing := invoiceCosting{
		Product:      "Starter",
		ProductPrice: 20.00,
	}

	costing.InvoiceSubtotal = costing.ProductPrice
	costing.InvoiceTax = costing.InvoiceSubtotal * 0.0875
	costing.InvoiceTotal = costing.InvoiceSubtotal + costing.InvoiceTax

	payment := paymentDetails{
		CCNumber:                   5634,
		BillingDetailsName:         "Adomate LLC",
		BillingDetailsAddressLine1: "17350 State Highway 249 STE 220",
		BillingDetailsAddressLine2: "Houston",
		BillingDetailsAddressState: "TX",
		BillingDetailsAddressZip:   77064,
	}

	data := emailData{
		InvoiceData:    invoice,
		InvoiceCosting: costing,
		PaymentDetails: payment,
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
