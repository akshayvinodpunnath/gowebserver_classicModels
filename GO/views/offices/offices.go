package offices

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akshayvinodpunnath/webserver/db"
	"github.com/gorilla/mux"

	"github.com/akshayvinodpunnath/webserver/models"
)

func GetOffices() []dbModels.Office {
	var office []dbModels.Office
	db := db.DbConn()
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM offices ORDER BY officeCode DESC")

	defer rows.Close()

	for rows.Next() {
		o := new(dbModels.Office)
		rows.Scan(&o.OfficeCode, &o.City, &o.Phone, &o.AddressLine1, &o.AddressLine2, &o.State, &o.Country, &o.PostalCode, &o.Territory)
		office = append(office, *o)
	}
	//fmt.Println(employee)
	return office
}

func GetOffice(OfficeCode string) []dbModels.Office {
	var office []dbModels.Office
	db := db.DbConn()
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM offices where OfficeCode = ?", OfficeCode)

	defer rows.Close()

	for rows.Next() {
		o := new(dbModels.Office)
		rows.Scan(&o.OfficeCode, &o.City, &o.Phone, &o.AddressLine1, &o.AddressLine2, &o.State, &o.Country, &o.PostalCode, &o.Territory)
		office = append(office, *o)
	}
	//fmt.Println(employee)
	return office
}

func OfficeMethods(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var office []dbModels.Office
		office = GetOffices()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(office)
	default:
		w.WriteHeader(http.StatusPreconditionFailed)
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func OfficeDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var office []dbModels.Office
	office = GetOffice(params["officeCode"])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(office)
}
