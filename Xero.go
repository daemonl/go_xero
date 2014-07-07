package xero

import (
	"encoding/json"
	"log"
	"strings"
)

type Xero struct {
	OAuth OAuth
}

type OAuth interface {
	DoGET(apiPath string) ([]byte, int, error)
	DoPOST(apiPath string, body []byte) ([]byte, int, error)
	DoPUT(apiPath string, body []byte) ([]byte, int, error)
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

func (x *Xero) Get(objectType string, pathParameters ...string) (interface{}, error) {
	apiPath := preparePath(objectType, pathParameters...)
	rawBytes, status, err := x.OAuth.DoGET(apiPath)
	if err != nil {
		return nil, err
	}

	return genericResponse(rawBytes, status)
}

func (x *Xero) Post(objectType string, obj interface{}, pathParameters ...string) (interface{}, error) {

	// Prepare {invoices:[{...}]}
	requestObject := map[string]interface{}{}

	objectType = strings.Title(objectType)
	requestObject[objectType] = []interface{}{
		obj,
	}

	body, err := json.Marshal(requestObject)
	if err != nil {
		return nil, err
	}

	apiPath := preparePath(objectType, pathParameters...)

	log.Printf("XERO POST %s\n%s\n", apiPath, string(body))
	resp, status, err := x.OAuth.DoPOST(apiPath, body)
	if err != nil {
		log.Printf("XERO Error Response\nRESP: %s\nERR: %s\n", string(resp), err.Error())
		return nil, err
	}

	log.Printf("XERO Response: %s\n", string(resp))

	return genericResponse(resp, status)
}

func preparePath(objectType string, pathParameters ...string) string {

	apiPath := objectType
	for i, param := range pathParameters {
		if i > 0 && param[0:1] != "?" {
			apiPath += "/"
		}
		apiPath += param
	}
	return apiPath
}

func genericResponse(resp []byte, status int) (interface{}, error) {

	switch status {
	case 200: // OK
		var dest interface{}
		err := json.Unmarshal(resp, &dest)
		return resp, err

	case 400: // Bad Request
		var dest APIException
		err := json.Unmarshal(resp, &dest)
		if err != nil {
			return nil, err
		}
		dest.HTTPStatus = status
		// Return as an error
		return nil, &dest

	case 401:
		return nil, &APIException{
			HTTPStatus: 401,
			Message:    "Invalid Authorization Credentials",
			Type:       "HTTP",
		}

	case 403:
		return nil, &APIException{
			HTTPStatus: 403,
			Message:    "The client SSL certificate was not valid.",
			Type:       "HTTP",
		}

	case 404:
		return nil, &APIException{
			HTTPStatus: 404,
			Message:    "Not Found",
			Type:       "HTTP",
		}

	case 500:
		return nil, &APIException{
			HTTPStatus: 500,
			Message:    "Xero Server Error",
			Type:       "HTTP",
		}

	case 501:
		return nil, &APIException{
			HTTPStatus: 501,
			Message:    "Not Implemented",
			Type:       "HTTP",
		}
	case 503:
		return nil, &APIException{
			HTTPStatus: 503,
			Message:    "Rate Limited or service unavailable",
			Type:       "HTTP",
		}
	}

	return nil, &APIException{
		HTTPStatus: status,
		Message:    "Unknown Status Code",
		Type:       "HTTP",
	}

}

/*
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
*/
