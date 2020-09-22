package routes

import (
	"log"
	"net/http"

	"github.com/akshayvinodpunnath/webserver/views/customers"
	"github.com/gorilla/mux"

	"github.com/akshayvinodpunnath/webserver/views/employees"

	"github.com/akshayvinodpunnath/webserver/views/offices"

	"fmt"
)

func UrlRoutes() {

	router := mux.NewRouter()
	fmt.Println("Webserver started")
	router.HandleFunc("/employees", employees.EmployeeMethods).Methods("GET")
	router.HandleFunc("/customers", customers.CustomerMethods).Methods("GET")
	router.HandleFunc("/offices", offices.OfficeMethods).Methods("GET")
	router.HandleFunc("/offices/{officeCode}", offices.OfficeDetails).Methods("GET")

	/*

		http.HandleFunc("/employees", employees.EmployeeMethods)
		http.HandleFunc("/customers", customers.CustomerMethods)
		http.HandleFunc("/offices", offices.OfficeMethods)
	*/
	//http.HandleFunc("/offices/{OfficeCode}", offices.OfficeDetails)
	log.Fatal(http.ListenAndServe(":8080", router))
}
