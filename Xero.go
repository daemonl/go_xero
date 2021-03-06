package xero

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"

	"github.com/daemonl/go_xero/xero_objects"
)

type Xero struct {
	OAuth OAuth
}

const (
	NormalApiPath  string = "api.xro/2.0"
	PayrollApiPath string = "payroll.xro/1.0"
)

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

func (x *Xero) Get(pathBase string, objectType string, pathParameters ...string) (interface{}, error) {
	apiPath := preparePath(objectType, pathParameters...)
	rawBytes, status, err := x.OAuth.DoGET(pathBase + "/" + apiPath)
	if err != nil {
		return nil, err
	}

	return genericResponse(rawBytes, status)
}

func (x *Xero) Post(obj interface{}, pathParameters ...string) (interface{}, error) {

	// Prepare {invoices:[{...}]}
	var requestObject interface{}
	var responseObject interface{}
	var objectType string
	var basePath string

	switch obj := obj.(type) {
	case *xero_objects.Invoice:
		requestObject = xero_objects.InvoiceRequest{
			Invoices: []*xero_objects.Invoice{
				obj,
			},
		}
		basePath = NormalApiPath
		objectType = "Invoices"
		responseObject = &xero_objects.InvoiceResponse{}

	case *xero_objects.LeaveApplication:
		requestObject = xero_objects.LeaveApplicationRequest{
			LeaveApplications: []*xero_objects.LeaveApplication{
				obj,
			},
		}
		basePath = PayrollApiPath
		objectType = "LeaveApplications"
		//responseObject = &xero_objects.LeaveApplicationResponse{}

	case *xero_objects.Timesheet:
		requestObject = xero_objects.TimesheetRequest{
			Timesheets: []*xero_objects.Timesheet{
				obj,
			},
		}
		basePath = PayrollApiPath
		objectType = "Timesheets"
		responseObject = &xero_objects.TimesheetResponse{}

	default:
		log.Println("No handler exists for %T", obj)
		return nil, fmt.Errorf("Object type %T is not yet implemented", obj)
	}

	body, err := xml.Marshal(requestObject)
	if err != nil {
		log.Printf("XML Marshal: %s\n", err.Error)
		return nil, err
	}
	apiPath := preparePath(objectType, pathParameters...)
	log.Printf("XERO POST %s\n%s\n", apiPath, string(body))

	resp, status, err := x.OAuth.DoPOST(basePath+"/"+apiPath, body)
	if err != nil {
		log.Printf("XERO Error Response\nRESP: %s\nERR: %s\n", string(resp), err.Error())
		return nil, err
	}

	log.Printf("XERO Response Raw: %s\n", string(resp))

	if status == 200 {
		log.Printf("XERO OK\n")
		err := xml.Unmarshal(resp, responseObject)
		if err == nil {
			return responseObject, nil
		}
		log.Printf("XERO Unmarshal error: %s\n", err.Error())
	}

	return genericResponse(resp, status)
}

type XeroObjectType string

const (
	XeroObjectType_Invoices          = "Invoices"
	XeroObjectType_LeaveApplications = "LeaveApplications"
	XeroObjectType_Timesheets        = "Timesheets"
)

func (x *Xero) GetObject(objectType XeroObjectType, pathParameters ...string) (interface{}, error) {

	var responseObject interface{}
	var basePath string
	switch objectType {
	case XeroObjectType_Invoices:
		basePath = NormalApiPath
		responseObject = &xero_objects.InvoiceResponse{}

	case XeroObjectType_LeaveApplications:
		basePath = PayrollApiPath
		//responseObject = &xero_objects.LeaveApplicationResponse{}

	case XeroObjectType_Timesheets:
		basePath = PayrollApiPath
		responseObject = &xero_objects.TimesheetResponse{}

	default:
		log.Println("No handler exists for %s", objectType)
		return nil, fmt.Errorf("Object type %s is not yet implemented", objectType)
	}

	apiPath := preparePath(string(objectType), pathParameters...)
	resp, status, err := x.OAuth.DoGET(basePath + "/" + apiPath)
	if err != nil {
		return nil, err
	}

	if status == 200 {
		log.Printf("XERO OK\n")
		err := xml.Unmarshal(resp, responseObject)
		if err == nil {
			return responseObject, nil
		}
		log.Printf("XERO Unmarshal error: %s\n", err.Error())
	}

	return genericResponse(resp, status)
}

func preparePath(objectType string, pathParameters ...string) string {
	if len(pathParameters) < 1 {
		return objectType
	}
	return objectType + "/" + strings.Join(pathParameters, "/")
}

func genericResponse(resp []byte, status int) (interface{}, error) {

	log.Printf("Generic response, status: %d\n", status)
	switch status {
	case 200: // OK
		var dest interface{}
		err := xml.Unmarshal(resp, &dest)
		if err != nil {
			fmt.Printf("Error unmarshalling xero response: %s\n", err.Error())
			fmt.Println(string(resp))
		}

		return string(resp), err

	case 400: // Bad Request
		var dest APIException
		err := xml.Unmarshal(resp, &dest)
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
