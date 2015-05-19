package xero_objects

type AddressType string

const (
	AddressType_POBOX    AddressType = "POBOX"
	AddressType_STREET   AddressType = "STREET"
	AddressType_DELIVERY AddressType = "DELIVERY"
)

type Address struct {
	Type        AddressType `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	Line1       string      `json:"AddressLine1,omitempty" xml:"AddressLine1,omitempty"`
	Line2       string      `json:"AddressLine2,omitempty" xml:"AddressLine2,omitempty"`
	Line3       string      `json:"AddressLine3,omitempty" xml:"AddressLine3,omitempty"`
	Line4       string      `json:"AddressLine4,omitempty" xml:"AddressLine4,omitempty"`
	City        string      `json:"City,omitempty" xml:"City,omitempty"`
	Region      string      `json:"Region,omitempty" xml:"Region,omitempty"`
	PostalCode  string      `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	Country     string      `json:"Country,omitempty" xml:"Country,omitempty"`
	AttentionTo string      `json:"AttentionTo,omitempty" xml:"AttentionTo,omitempty"`
}