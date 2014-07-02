package xero_objects

type Contact struct {
	ContactID     string `json:"ContactID,omitempty"` // Xero's ID
	ContactNumber string `json:"ContactNumber"`       // Our ID
	Name          string `json:"Name"`
}
