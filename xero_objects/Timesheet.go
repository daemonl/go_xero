package xero_objects

import "encoding/xml"

type TimesheetStatus string

const (
	TimesheetStatus_DRAFT     TimesheetStatus = "DRAFT"
	TimesheetStatus_PROCESSED TimesheetStatus = "PROCESSED"
	TimesheetStatus_APPROVED  TimesheetStatus = "APPROVED"
)

type Timesheet struct {
	TimesheetID string `xml:"TimesheetID,omitempty"`
	EmployeeID  string `xml:"EmployeeID"`
	StartDate   string `xml:"StartDate"`
	EndDate     string `xml:"EndDate"`

	Status TimesheetStatus `xml:"Status,omitempty"`
	Hours  float64         `xml:"Hours,omitempty"`

	TimesheetLines []TimesheetLine `xml:"TimesheetLines>TimesheetLine"`
}

type TimesheetLine struct {
	EarningsRateID string    `xml:"EarningsRateID"`
	Units          []float64 `xml:"NumberOfUnits>NumberOfUnit"`
	UpdatedDateUTC string    `xml:"UpdatedDateUTC,omitempty"`
}

type TimesheetRequest struct {
	XMLName    xml.Name     `xml:"Timesheets"`
	Timesheets []*Timesheet `xml:"Timesheet"`
}

type TimesheetResponse struct {
	Timesheets []Timesheet `xml:"Timesheets"`
}
