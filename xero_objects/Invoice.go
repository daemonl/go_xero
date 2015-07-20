package xero_objects

import (
	"encoding/xml"
)

type InvoiceType string

const (
	InvoiceType_ACCPAY InvoiceType = "ACCPAY" // A bill – commonly known as a Accounts Payable or supplier invoice
	InvoiceType_ACCREC InvoiceType = "ACCREC" // A sales invoice – commonly known as an Accounts Receivable or customer invoice
)

type InvoiceStatus string

const (
	InvoiceStatus_DRAFT      InvoiceStatus = "DRAFT"      // A Draft Invoice (default)
	InvoiceStatus_SUBMITTED  InvoiceStatus = "SUBMITTED"  // An Awaiting Approval Invoice
	InvoiceStatus_DELETED    InvoiceStatus = "DELETED"    // A Deleted Invoice
	InvoiceStatus_AUTHORISED InvoiceStatus = "AUTHORISED" // An Invoice that is Approved and Awaiting Payment OR partially paid
	InvoiceStatus_PAID       InvoiceStatus = "PAID"       // An Invoice that is completely Paid
	InvoiceStatus_VOIDED     InvoiceStatus = "VOIDED"     // A Voided Invoice
)

type InvoiceLineAmountType string

const (
	InvoiceLineAmountType_EXCLUSIVE InvoiceLineAmountType = "Exclusive" // Invoice lines are exclusive of tax (default)
	InvoiceLineAmountType_INCLUSIVE InvoiceLineAmountType = "Inclusive" // Invoice lines are inclusive tax
	InvoiceLineAmountType_NOTAX     InvoiceLineAmountType = "NoTax"     // Invoices lines have no tax
)

type InvoiceLineTaxType string

const (
	InvoiceLineTaxType_OUTPUT          InvoiceLineTaxType = "OUTPUT"          //	10.00	GST on Income
	InvoiceLineTaxType_INPUT           InvoiceLineTaxType = "INPUT"           //	10.00	GST on Expenses
	InvoiceLineTaxType_CAPEXINPUT      InvoiceLineTaxType = "CAPEXINPUT"      //	10.00	GST on Capital	Yes
	InvoiceLineTaxType_EXEMPTEXPORT    InvoiceLineTaxType = "EXEMPTEXPORT"    //	0.00	GST Free Exports	Yes
	InvoiceLineTaxType_EXEMPTEXPENSES  InvoiceLineTaxType = "EXEMPTEXPENSES"  //	0.00	GST Free Expenses	Yes
	InvoiceLineTaxType_EXEMPTCAPITAL   InvoiceLineTaxType = "EXEMPTCAPITAL"   //	0.00	GST Free Capital	Yes
	InvoiceLineTaxType_EXEMPTOUTPUT    InvoiceLineTaxType = "EXEMPTOUTPUT"    //	0.00	GST Free Income
	InvoiceLineTaxType_INPUTTAXED      InvoiceLineTaxType = "INPUTTAXED"      //	0.00	Input Taxed
	InvoiceLineTaxType_BASEXCLUDED     InvoiceLineTaxType = "BASEXCLUDED"     //	0.00	BAS Excluded
	InvoiceLineTaxType_GSTONCAPIMPORTS InvoiceLineTaxType = "GSTONCAPIMPORTS" //	0.00	GST on Capital Imports	Yes
	InvoiceLineTaxType_GSTONIMPORTS    InvoiceLineTaxType = "GSTONIMPORTS"    //	0.00	GST on Imports	Yes
)

type InvoiceResponse struct {
	ID           string     `json:"Id"`
	Status       string     `json:"Status"`
	ProviderName string     `json:"ProviderName"`
	DateTimeUTC  string     `json:"DateTimeUTC"`
	Invoices     []*Invoice `json:"Invoices" xml:"Invoices>Invoice"`
}

type InvoiceRequest struct {
	XMLName  xml.Name   `xml:"Invoices"`
	Invoices []*Invoice `json:"Invoices" xml:"Invoice"`
}

type Invoice struct {
	ID                  string                `json:"InvoiceID,omitempty" xml:"InvoiceID,omitempty"`
	InvoiceNumber       string                `json:"InvoiceNumber,omitempty" xml:"InvoiceNumber,omitempty"`
	Type                InvoiceType           `json:"Type,omitempty" xml:"Type,omitempty"`
	Status              InvoiceStatus         `json:"Status,omitempty" xml:"Status,omitempty"`
	Contact             Contact               `json:"Contact,omitempty" xml:"Contact,omitempty"`
	Date                string                `json:"Date,omitempty" xml:"Date,omitempty"`
	DueDate             string                `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	ExpectedPaymentDate string                `json:"ExpectedPaymentDate,omitempty" xml:"ExpectedPaymentDate,omitempty"`
	Reference           string                `json:"Reference,omitempty" xml:"Reference,omitempty"`
	BrandingThemeID     string                `json:"BrandingThemeID,omitempty" xml:"BrandingThemeID,omitempty"`
	URL                 string                `json:"Url,omitempty" xml:"Url,omitempty"`
	CurrencyCode        string                `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`
	LineAmountTypes     InvoiceLineAmountType `json:"LineAmountType,omitempty" xml:"LineAmountType,omitempty"`
	SubTotal            *float64              `json:"SubTotal,omitempty" xml:"SubTotal,omitempty"`
	Total               *float64              `json:"Total,omitempty" xml:"Total,omitempty"`
	TotalTax            *float64              `json:"TotalTax,omitempty" xml:"TotalTax,omitempty"`
	LineItems           []InvoiceLineItem     `json:"LineItems" xml:"LineItems>LineItem"`
	SentToContact       bool                  `json:"SentToContact" xml:"SentToContact"`
}

type InvoiceLineItem struct {
	Description string  `json:"Description" xml:"Description"`
	Quantity    float64 `json:"Quantity" xml:"Quantity"`
	UnitAmount  float64 `json:"UnitAmount" xml:"UnitAmount"`
	ItemCode    string  `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
	AccountCode string  `json:"AccountCode" xml:"AccountCode"`

	TaxType      *InvoiceLineTaxType `json:"TaxType,omitempty" xml:"TaxType,omitempty"`
	TaxAmount    *float64            `json:"TaxAmount,omitempty" xml:"TaxAmount,omitempty"`
	LineAmount   *float64            `json:"LineAmount,omitempty" xml:"LineAmount,omitempty"`
	DiscountRate *float64            `json:"DiscountRate,omitempty" xml:"DiscountRate,omitempty"`
}
