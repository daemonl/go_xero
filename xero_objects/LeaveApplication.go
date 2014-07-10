package xero_objects

import (
	"encoding/xml"
)

type LeaveApplicationRequest struct {
	XMLName           xml.Name            `xml:"LeaveApplications"`
	LeaveApplications []*LeaveApplication `json:"LeaveApplications" xml:"LeaveApplication"`
}
type LeaveApplication struct {
	EmployeeID   string        `json:"EmployeeID" xml:"EmployeeID"`
	LeaveTypeID  string        `json:"LeaveTypeID" xml:"LeaveTypeID"`
	Title        string        `json:"Title" xml:"Title"`
	StartDate    string        `json:"StartDate" xml:"StartDate"`
	EndDate      string        `json:"EndDate" xml:"EndDate"`
	Description  string        `json:"Description" xml:"Description"`
	LeavePeriods []LeavePeriod `json:"LeavePeriods" xml:"LeavePeriods>LeavePeriod"`
}

type LeavePeriod struct {
	PayPeriodEndDate string  `json:"PayPeriodEndDate" xml:"PayPeriodEndDate"`
	NumberOfUnits    float64 `json:"NumberOfUnits" xml:"NumberOfUnits"`
}
