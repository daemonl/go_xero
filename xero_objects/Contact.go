package xero_objects

type Contact struct {
	ContactID string `json:"ContactID,omitempty" xml:"ContactID,omitempty"`
	ContactNumber string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	ContactStatus string `json:"ContactStatus,omitempty" xml:"ContactStatus,omitempty"`
	Name string `json:"Name" xml:"Name"`
	FirstName string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	LastName string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	EmailAddress string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`
	SkypeUserName string `json:"SkypeUserName,omitempty" xml:"SkypeUserName,omitempty"`
	BankAccountDetails string `json:"BankAccountDetails,omitempty" xml:"BankAccountDetails,omitempty"`
	TaxNumber string `json:"TaxNumber,omitempty" xml:"TaxNumber,omitempty"`
	AccountsReceivableTaxType string `json:"AccountsReceivableTaxType,omitempty" xml:"AccountsReceivableTaxType,omitempty"`
	AccountsPayableTaxType string `json:"AccountsPayableTaxType,omitempty" xml:"AccountsPayableTaxType,omitempty"`
	Addresses []Address `json:"Addresses,omitempty" xml:"Addresses,omitempty"`
	Phones []Phone `json:"Phones,omitempty" xml:"Phones,omitempty"`
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty" xml:"UpdatedDateUTC,omitempty"`
	IsSupplier bool `json:"IsSupplier,omitempty" xml:"IsSupplier,omitempty"`
	IsCustomer bool `json:"IsCustomer,omitempty" xml:"IsCustomer,omitempty"`
	DefaultCurrency string `json:"DefaultCurrency,omitempty" xml:"DefaultCurrency,omitempty"`
	XeroNetworkKey string `json:"XeroNetworkKey,omitempty" xml:"XeroNetworkKey,omitempty"`
}

type AddressType string

const (
	AddressType_POBOX    = "POBOX"
	AddressType_STREET   = "STREET"
	AddressType_DELIVERY = "DELIVERY"
)

type Address struct {
	Type        AddressType `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	Line1       string      `json:"AddressLine1,omitempty" xml:"AddressLine1,omitempty"`
	Line2       string      `json:"AddressLine1,omitempty" xml:"AddressLine1,omitempty"`
	Line3       string      `json:"AddressLine2,omitempty" xml:"AddressLine2,omitempty"`
	Line4       string      `json:"AddressLine3,omitempty" xml:"AddressLine3,omitempty"`
	City        string      `json:"City,omitempty" xml:"City,omitempty"`
	Region      string      `json:"Region,omitempty" xml:"Region,omitempty"`
	PostalCode  string      `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	Country     string      `json:"Country,omitempty" xml:"Country,omitempty"`
	AttentionTo string      `json:"AttentionTo,omitempty" xml:"AttentionTo,omitempty"`
}

type PhoneType string

const (
	PhoneType_DEFAULT = "DEFAULT"
	PhoneType_DDI     = "DDI"
	PhoneType_MOBILE  = "MOBILE"
	PhoneType_FAX     = "FAX"
)

type Phone struct {
	Type        PhoneType `json":" xml:":"PhoneType,omitempty"`
	Number      string    `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	AreaCode    string    `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	CountryCode string    `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
}
