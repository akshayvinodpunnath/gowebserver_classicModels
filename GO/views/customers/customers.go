package customers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akshayvinodpunnath/webserver/db"

	"github.com/akshayvinodpunnath/webserver/models"
)

func GetCustomers() []dbModels.Customer {
	var customer []dbModels.Customer
	db := db.DbConn()
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM customers ORDER BY customerNumber DESC")

	defer rows.Close()

	for rows.Next() {
		c := new(dbModels.Customer)
		rows.Scan(&c.CustomerNumber, &c.CustomerName, &c.ContactLastName, &c.ContactFirstName, &c.Phone, &c.AddressLine1, &c.AddressLine2, &c.City, &c.State, &c.PostalCode, &c.Country, &c.SalesRepEmployeeNumber, &c.CreditLimit)

		customer = append(customer, *c)
	}
	//fmt.Println(customer)
	return customer
}

func CustomerMethods(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var customer []dbModels.Customer
		customer = GetCustomers()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	default:
		w.WriteHeader(http.StatusPreconditionFailed)
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
