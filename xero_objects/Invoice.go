package xero_objects

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
	Invoices     []*Invoice `json:"Invoices"`
}

type InvoiceRequest struct {
	Invoices []*Invoice `json:"Invoices"`
}

type Invoice struct {
	ID                  string                `json:"InvoiceID,omitempty"`
	InvoiceNumber       string                `json:"InvoiceNumber,omitempty"`
	Type                InvoiceType           `json:"Type,omitempty"`
	Status              InvoiceStatus        `json:"Status,omitempty"`
	Contact             Contact              `json:"Contact,omitempty"`
	Date                string               `json:"Date,omitempty"`
	DueDate             string               `json:"DueDate,omitempty"`
	ExpectedPaymentDate string               `json:"ExpectedPaymentDate,omitempty"`
	Reference           string               `json:"Reference,omitempty"`
	BrandingThemeID     string               `json:"BrandingThemeID,omitempty"`
	URL                 string               `json:"Url,omitempty"`
	CurrencyCode        string               `json:"CurrencyCode,omitempty"`
	LineAmountTypes     InvoiceLineAmountType `json:"LineAmountType,omitempty"`
	SubTotal            *float64              `json:"SubTotal,omitempty"`
	Total               *float64              `json:"Total,omitempty"`
	TotalTax            *float64              `json:"TotalTax,omitempty"`
	LineItems           []InvoiceLineItem     `json:"LineItems"`
}

type InvoiceLineItem struct {
	Description string  `json:"Description"`
	Quantity    float64 `json:"Quantity"`
	UnitAmount  float64 `json:"UnitAmount"`
	ItemCode    string  `json:"ItemCode,omitempty"`
	AccountCode string  `json:"AccountCode"`

	TaxType      *InvoiceLineTaxType `json:"TaxType,omitempty"`
	TaxAmount    *float64            `json:"TaxAmount,omitempty"`
	LineAmount   *float64            `json:"LineAmount,omitempty"`
	DiscountRate *float64            `json:"DiscountRate,omitempty"`
}
