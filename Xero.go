package xero

import (
	"encoding/json"
	"fmt"

	"github.com/daemonl/go_xero/xero_objects"
)

type Xero struct {
	OAuth OAuth
}

type OAuth interface {
	DoGET(apiPath string) ([]byte, error)
	DoPOST(apiPath string, body []byte) ([]byte, error)
	DoPUT(apiPath string, body []byte) ([]byte, error)
}

func GetXeroPrivateCore(keyFilename string, key string) (*Xero, error) {
	privateKey, err := loadPrivateKeyFromFile(keyFilename)
	if err != nil {
		return nil, err
	}

	oac := &xeroPrivateOAuth{
		privateKey:  privateKey,
		consumerKey: key,
	}

	xero := &Xero{
		OAuth: oac,
	}
	return xero, nil
}

func (x *Xero) GetToObject(apiPath string, dest interface{}) error {
	rawBytes, err := x.OAuth.DoGET(apiPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawBytes, dest)
	if err != nil {
		return err
	}
	return nil
}

func (x *Xero) PostObject(apiPath string, obj interface{}, dest interface{}) error {
	body, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	resp, err := x.OAuth.DoPOST(apiPath, body)
	fmt.Println(string(resp))

	err = json.Unmarshal(resp, dest)
	if err != nil {
		return err
	}
	return nil
}

func (x *Xero) GetInvoice(invoiceID string) (*xero_objects.Invoice, error) {
	response := &xero_objects.InvoiceResponse{}
	err := x.GetToObject("Invoice/"+invoiceID, response)
	if err != nil {
		return nil, err
	}
	if len(response.Invoices) < 1 {
		return nil, nil
	}
	return response.Invoices[0], nil
}

func (x *Xero) PostInvoice(invoice *xero_objects.Invoice) (*xero_objects.Invoice, error) {
	req := xero_objects.InvoiceRequest{
		Invoices: []*xero_objects.Invoice{
			invoice,
		},
	}
	response := &xero_objects.InvoiceResponse{}
	err := x.PostObject("Invoice", req, response)
	if err != nil {
		return nil, err
	}
	if response.Status != "OK" {
		return nil, fmt.Errorf("Error with the invoice")
	}
	if len(response.Invoices) < 1 {
		return nil, fmt.Errorf("No invoices returned")
	}
	return response.Invoices[0], nil
}
