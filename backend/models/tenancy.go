package models

type Tenancy struct {
	ID           string `json:"id"`
	LandlordName string `json:"landlord_name"`
	TenantName   string `json:"tenant_name"`
	PropertyID   string `json:"property_id"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	RentAmount   int    `json:"rent_amount"`
	Deposit      int    `json:"deposit"`
}
