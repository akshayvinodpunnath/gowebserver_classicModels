package dbModels

//import "github.com/akshayvinodpunnath/webserver/models/dataTypes"

//NullInt64 & NullString defined in dataTypes.go
type Employee struct {
	EmployeeNumber int       `json:"employeeNumber"`
	LastName       string    `json:"lastName"`
	FirstName      string    `json:"firstName"`
	Extension      string    `json:"extension"`
	Email          string    `json:"email"`
	OfficeCode     string    `json:"OfficeCode"`
	ReportsTo      NullInt64 `json:"reportsTo"`
	JobTitle       string    `json:"jobTitle"`
}

type Customer struct {
	CustomerNumber         int        `json:"customerNumber"`
	CustomerName           string     `json:"customerName"`
	ContactLastName        string     `json:"contactLastName"`
	ContactFirstName       string     `json:"contactFirstName"`
	Phone                  string     `json:"phone"`
	AddressLine1           string     `json:"addressLine1"`
	AddressLine2           NullString `json:"addressLine2"`
	City                   string     `json:"city"`
	State                  NullString `json:"state"`
	PostalCode             string     `json:"postalCode"`
	Country                string     `json:"country"`
	SalesRepEmployeeNumber int        `json:"salesRepEmployeeNumber"`
	CreditLimit            float64    `json:"creditLimit"`
}

type Office struct {
	OfficeCode   string     `json:"officeCode"`
	City         string     `json:"city"`
	Phone        string     `json:"phone"`
	AddressLine1 string     `json:"addressLine1"`
	AddressLine2 NullString `json:"addressLine2"`
	State        NullString `json:"state"`
	Country      string     `json:"country"`
	PostalCode   string     `json:"postalCode"`
	Territory    string     `json:"territory"`
}
