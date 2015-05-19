package xero_objects

type Employee struct {
	HomeAddress                Address           `xml:"HomeAddress,omitempty"`
	TaxDeclaration             TaxDeclaration    `xml:"TaxDeclaration,omitempty"`
	BankAccounts               []BankAccount     `xml:"BankAccounts>BankAccount,omitempty"`
	OpeningBalances            []OpeningBalance  `xml:"OpeningBalances,omitempty"`
	PayTemplates               []PayTemplate     `xml:"PayTemplate,omitempty"`
	SuperMemberships           []SuperMembership `xml:"SuperMemberships>SuperMembership,omitempty"`
	EmployeeID                 string            `xml:"EmployeeID,omitempty"`
	FirstName                  string            `xml:"FirstName"`
	LastName                   string            `xml:"LastName"`
	Status                     string            `xml:"Status,omitempty"`
	Email                      string            `xml:"Email"`
	DateOfBirth                string            `xml:"DateOfBirth"`
	Gender                     string            `xml:"Gender"`
	Phone                      string            `xml:"Phone"`
	Mobile                     string            `xml:"Mobile"`
	StartDate                  string            `xml:"StartDate,omitempty"`
	OrdinaryEarningsRateID     string            `xml:"OrdinaryEarningsRateID,omitempty"`
	PayrollCalendarID          string            `xml:"PayrollCalendarID,omitempty"`
	UpdatedDateUTC             string            `xml:"UpdatedDateUTC,omitempty"`
	IsAuthorisedToApproveLeave bool              `xml:"IsAuthorisedToApproveLeave,omitempty"`
}

type TaxDeclaration struct {
	TFNPendingOrExemptionHeld        bool   `xml:"TFNPendingOrExemptionHeld,omitempty"`
	AustralianResidentForTaxPurposes bool   `xml:"AustralianResidentForTaxPurposes,omitempty"`
	TaxFreeThresholdClaimed          bool   `xml:"TaxFreeThresholdClaimed,omitempty"`
	HasHELPDebt                      bool   `xml:"HasHELPDebt,omitempty"`
	HasSFSSDebt                      bool   `xml:"HasSFSSDebt,omitempty"`
	EligibleToReceiveLeaveLoading    bool   `xml:"EligibleToReceiveLeaveLoading,omitempty"`
	UpdatedDateUTC                   string `xml:"UpdatedDateUTC,omitempty"`
}

type BankAccount struct {
	StatementText string  `xml:"StatementText,omitempty"`
	AccountName   string  `xml:"AccountName,omitempty"`
	BSB           string  `xml:"BSB,omitempty"`
	AccountNumber string  `xml:"AccountNumber,omitempty"`
	Remainder     bool    `xml:"Remainder,omitempty"`
	Percentage    float64 `xml:"Percentage,omitempty"`
}

type OpeningBalance struct {
	EarningsLines      []EarningsLine      `xml:"EarningsLines>EarningsLine,omitempty"`
	DeductionLines     []DeductionLine     `xml:"DeductionLines>DeductionLine,omitempty"`
	SuperLines         []SuperLine         `xml:"SuperLines>SuperLine,omitempty"`
	ReimbursementLines []ReimbursementLine `xml:"ReimbursementLines>ReimbursementLine,omitempty"`
	LeaveLines         []LeaveLine         `xml:"LeaveLines>LeaveLine,omitempty"`
	Tax                float64             `xml:"Tax,omitempty"`
	OpeningBalanceDate string              `xml:"OpeningBalanceDate,omitempty"`
}

type PayTemplate struct {
	EarningsLines      []EarningsLine      `xml:"EarningsLines>EarningsLine,omitempty"`
	DeductionLines     []DeductionLine     `xml:"DeductionLines>DeductionLine,omitempty"`
	SuperLines         []SuperLine         `xml:"SuperLines>SuperLine,omitempty"`
	ReimbursementLines []ReimbursementLine `xml:"ReimbursementLines>ReimbursementLine,omitempty"`
	LeaveLines         []LeaveLine         `xml:"LeaveLines>LeaveLine,omitempty"`
}

type EarningsLine struct {
	EarningsRateID       string  `xml:"EarningsRateID,omitempty"`
	Amount               float64 `xml:"Amount,omitempty"`
	CalculationType      string  `xml:"CalculationType,omitempty"`
	AnnualSalary         float64 `xml:"AnnualSalary,omitempty"`
	NumberOfUnitsPerWeek float64 `xml:"NumberOfUnitsPerWeek,omitempty"`
}

type DeductionLine struct {
	DeductionTypeID string  `xml:"DeductionTypeID,omitempty"`
	Amount          float64 `xml:"Amount,omitempty"`
}

type SuperLine struct {
	SuperMembershipID      string  `xml:"SuperMembershipID,omitempty"`
	ContributionType       string  `xml:"ContributionType,omitempty"`
	Amount                 float64 `xml:"Amount,omitempty"`
	CalculationType        string  `xml:"CalculationType,omitempty"`
	MinimumMonthlyEarnings float64 `xml:"MinimumMonthlyEarnings,omitempty"`
	ExpenseAccountCode     int     `xml:"ExpenseAccountCode,omitempty"`
	LiabilityAccountCode   int     `xml:"LiabilityAccountCode,omitempty"`
	Percentage             float64 `xml:"Percentage,omitempty"`
}

type ReimbursementLine struct {
	ReimbursementTypeID string  `xml:"ReimbursementTypeID,omitempty"`
	Description			string	`xml:"Description,omitempty"`
	Amount              float64 `xml:"Amount,omitempty"`
}

type LeaveLine struct {
	LeaveTypeID         string  `xml:"LeaveTypeID,omitempty"`
	NumberOfUnits       float64 `xml:"NumberOfUnits,omitempty"`
	CalculationType     string  `xml:"CalculationType,omitempty"`
	AnnualNumberOfUnits string  `xml:"AnnualNumberOfUnits,omitempty"`
}

type SuperMembership struct {
	SuperMembershipID string `xml:"SuperMembershipID,omitempty"`
	SuperFundID string `xml:"SuperFundID,omitempty"`
	EmployeeNumber    string `xml:"EmployeeNumber,omitempty"`
}
