package xero_objects

type Contact struct {
	//ContactPersons            string `json:"ContactPersons,omitempty"`
	ContactID string `json:"ContactID,omitempty"`
	//               Xero identifier
	ContactNumber string `json:"ContactNumber,omitempty"`
	//               This can be updated via the API only i.e. This field is read only on the Xero contact screen, used to identify contacts in external systems (max length = 50)
	ContactStatus string `json:"ContactStatus,omitempty"`
	// Current status of a contact – see contact status types
	Name string `json:"Name"`
	// Name of contact organisation (max length = 500)
	FirstName string `json:"FirstName,omitempty"`
	// First name of contact person (max length = 255)
	LastName string `json:"LastName,omitempty"`
	// Last name of contact person (max length = 255)
	EmailAddress string `json:"EmailAddress,omitempty"`
	// Email address of contact person (umlauts not supported) (max length = 500)
	SkypeUserName string `json:"SkypeUserName,omitempty"`
	//  Skype user name of contact
	BankAccountDetails string `json:"BankAccountDetails,omitempty"`
	// Bank account number of contact
	TaxNumber string `json:"TaxNumber,omitempty"`
	// Tax number of contact – this is also known as the ABN (Australia), GST Number (New Zealand), VAT Number (UK) or Tax ID Number (US and global) in the Xero UI depending on which regionalized version of Xero you are using (max length = 50)
	AccountsReceivableTaxType string `json:"AccountsReceivableTaxType,omitempty"`
	// Default tax type used for contact on AR invoices
	AccountsPayableTaxType string `json:"AccountsPayableTaxType,omitempty"`
	// Default tax type used for contact on AP invoices
	Addresses []Address `json:"Addresses,omitempty"`
	// Store certain address types for a contact – see address types
	Phones []Phone `json:"Phones,omitempty"`
	//  Store certain phone types for a contact – see phone types
	UpdatedDateUTC string `json:"UpdatedDateUTC,omitempty"`
	// UTC timestamp of last update to contact
	IsSupplier bool `json:"IsSupplier,omitempty"`
	//  READ ONLY Describes if a contact that has any AP invoices entered against them.
	IsCustomer bool `json:"IsCustomer,omitempty"`
	//  READ ONLY Describes if a contact has any AR invoices entered against them.
	DefaultCurrency string `json:"DefaultCurrency,omitempty"`
	//  Default currency for raising invoices against contact
	XeroNetworkKey string `json:"XeroNetworkKey,omitempty"`
	//  Store XeroNetworkKey for contacts. Optional element for PUT/POST requests. Retrieved via GET request for a specific contact or pagination
}

type AddressType string

const (
	AddressType_POBOX    = "POBOX"
	AddressType_STREET   = "STREET"
	AddressType_DELIVERY = "DELIVERY"
)

type Address struct {
	Type        AddressType `json:"AddressType,omitempty"`
	Line1       string      `json:"AddressLine1,omitempty"`
	Line2       string      `json:"AddressLine1,omitempty"`
	Line3       string      `json:"AddressLine2,omitempty"`
	Line4       string      `json:"AddressLine3,omitempty"`
	City        string      `json:City,omitempty"`
	Region      string      `json:"Region,omitempty"`
	PostalCode  string      `json:"PostalCode,omitempty"`
	Country     string      `json:"Country,omitempty"`
	AttentionTo string      `json:"AttentionTo,omitempty"`
}

type PhoneType string

const (
	PhoneType_DEFAULT = "DEFAULT"
	PhoneType_DDI     = "DDI"
	PhoneType_MOBILE  = "MOBILE"
	PhoneType_FAX     = "FAX"
)

type Phone struct {
	Type        PhoneType `json":"PhoneType,omitempty"`
	Number      string    `json:"PhoneNumber,omitempty"`
	AreaCode    string    `json:"AreaCode,omitempty"`
	CountryCode string    `json:"CountryCode,omitempty"`
}
