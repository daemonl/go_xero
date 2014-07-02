package xero

import (
	"fmt"
	"github.com/daemonl/go_xero/xero_objects"
	"strings"
	"testing"
)

var testKey string = "6FKW0MJPFOCFWCFNL3CDJEZBNYVP4I"
var testPrivateKeyFile string = "/home/daemonl/Projects/xero/xero_dev_key.pem"

func Test_GetKey(t *testing.T) {

	pk, err := loadPrivateKeyFromFile(testPrivateKeyFile)
	//pk, err := loadPrivateKeyFromFile("/home/daemonl/.ssh/id_rsa")
	if err != nil {
		t.Fatal(err.Error())
	}
	if pk == nil {
		t.Fatal("No private key loaded")
	}
}

func Test_request(t *testing.T) {

	xero, err := GetXeroPrivateCore(testPrivateKeyFile, testKey)
	if err != nil {
		t.Fatal(err.Error())
	}
	h, err := xero.OAuth.DoGET("Invoice/8b0ccb6a-d9b7-4da5-8360-ef7fb157b5aa")
	if err != nil {
		t.Fatal(err.Error())
	}
	hStr := string(h)
	if !strings.Contains(hStr, "Invoice") {
		t.Fail()
	}
}

func Test_invoice(t *testing.T) {
	xero, err := GetXeroPrivateCore(testPrivateKeyFile, testKey)
	if err != nil {
		t.Fatal(err.Error())
	}
	inv, err := xero.GetInvoice("8b0ccb6a-d9b7-4da5-8360-ef7fb157b5aa")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println("INVOICE")
	fmt.Println(inv)
}

func Test_post(t *testing.T) {
	xero, err := GetXeroPrivateCore(testPrivateKeyFile, testKey)
	if err != nil {
		t.Fatal(err.Error())
	}

	date := "2014-05-01T00:00:00"
	dateDue := "2014-06-01T00:00:00"

	inv := &xero_objects.Invoice{
		ID:   "a7fbffe9-1a73-4a61-a88b-cae39b09f9d4",
		Type: xero_objects.InvoiceType_ACCREC,
		Contact: &xero_objects.Contact{
			Name: "John Smith",
		},
		Date:            &date,
		DueDate:         &dateDue,
		LineAmountTypes: xero_objects.InvoiceLineAmountType_EXCLUSIVE,
		LineItems: []xero_objects.InvoiceLineItem{
			xero_objects.InvoiceLineItem{
				Description: "Phase 1 - Thingo",
				Quantity:    0.3,
				UnitAmount:  100,
				AccountCode: "200",
			},
		},
	}

	respInv, err := xero.PostInvoice(inv)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("Invoice ID: %s\n", respInv.ID)
}
