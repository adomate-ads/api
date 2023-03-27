/*
Potential Bugs: 

HTML and GO variable names don't match
Different variable naming conventions

Needs Fixing:

Write all the function data in one function, and pass that function to the http
*/

package main

    import (
        "log"
        "html/template"
        "net/http"
        "fmt"
        "os"
        "os/signal"
        "syscall"
    )


    type invoiceData struct{
        invoiceNumber int
        companyName string
    }

    type invoiceCosting struct {
        product       string
        productPrice  float32
        invoiceSubtotal float32
        invoiceTax      float32
        invoiceTotal    float32
    }

    type paymentDetails struct {
        ccNumber int
        billingDetailsname string
        billingDetailsaddressLine1 string
        billingDetailsaddressLine2 string
        billingDetailsaddressState string
        billingDetailsaddressZipcode int
    }

    type emailData struct {
        InvoiceData     invoiceData
        InvoiceCosting  invoiceCosting
        PaymentDetails  paymentDetails
    }
    // Create handler for invoiceData
    func invoiceDataHandler(w http.ResponseWriter, r *http.Request) {
        t, _ := template.ParseFiles("invoice.html")
        data:= invoiceData{
            invoiceNumber: 1934,
            companyName: "Adomate LLC",
        }
          // TODO remove later, for testing purpose only
        fmt.Println("Invoice No:", data.invoiceNumber)
        fmt.Println("Company Name:", data.companyName)
        t.Execute(w,data)
    }

    // Create handler for invoiceCosting 
    func invoiceCostingHandler(w http.ResponseWriter, r *http.Request) {
        t, _ := template.ParseFiles("invoice.html")
        data:= invoiceCosting{
            product: "Starter",
            productPrice: 20.00,
        }
        data.invoiceSubtotal = data.productPrice
        data.invoiceTax = data.invoiceSubtotal * 0.0875
        data.invoiceTotal = data.invoiceSubtotal + data.invoiceTax
        // TODO remove later, for testing purpose only
            fmt.Println("Product:", data.product)
            fmt.Println("ProductPrice:", data.productPrice)
            fmt.Println("InvoiceSubtotal:", data.invoiceSubtotal)
            fmt.Println("InvoiceTax:", data.invoiceTax)
            fmt.Println("InvoiceTotal:", data.invoiceTotal)
        t.Execute(w,data)
    }

    // Create handler for paymentDetails
    func paymentDetailsHandler(w http.ResponseWriter, r *http.Request) {
        t, _ := template.ParseFiles("invoice.html")
        data:= paymentDetails{
            ccNumber: 1934,
            billingDetailsname: "Adomate LLC",
            billingDetailsaddressLine1: "",
            billingDetailsaddressLine2:"",
            billingDetailsaddressState:"TX",
            billingDetailsaddressZipcode:77479,
        }
          // TODO remove later, for testing purpose only
        fmt.Println("billingDetailsname:", data.billingDetailsname)
        fmt.Println("ProductPrice:", data.billingDetailsaddressLine1)
        fmt.Println("InvoiceSubtotal:", data.billingDetailsaddressLine2)
        fmt.Println("InvoiceTax:", data.billingDetailsaddressState)
        fmt.Println("InvoiceTotal:", data.billingDetailsaddressZipcode)
        t.Execute(w,data)
    }
    // TODO Create handler for all
    func invoiceEmailPageHandler(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("invoice.html")
		data:= emailData{

		}
		t.Execute(w,data)
    }
    func main() {
        http.HandleFunc("/AdomateInvoice", invoiceEmailPageHandler) // TODO update to invoiceEmailPageHandler

        log.Printf("Updating Invoice...")
        fmt.Println("Server is now running. Press CTRL-C to exit.")
        http.ListenAndServe(":8080", nil)
        sc := make(chan os.Signal, 1)
        signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
    }
