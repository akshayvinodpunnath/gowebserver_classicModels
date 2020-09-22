package employees

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akshayvinodpunnath/webserver/db"

	"github.com/akshayvinodpunnath/webserver/models"
)

func GetEmployees() []dbModels.Employee {
	var employee []dbModels.Employee
	db := db.DbConn()
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM employees ORDER BY employeeNumber DESC")

	defer rows.Close()

	for rows.Next() {
		e := new(dbModels.Employee)
		//var fName, lName, ext string
		rows.Scan(&e.EmployeeNumber, &e.LastName, &e.FirstName, &e.Extension, &e.Email, &e.OfficeCode, &e.ReportsTo, &e.JobTitle)
		//, &lName, &fName, &ext) , &e.email, &e.officeCode, &e.reportsTo, &e.jobTitle)
		//employee = append(employee, Employee{Id: employeeNumber, LastName: lName}) //, lastName: lName, firstName: fName, extension: ext})
		employee = append(employee, *e)
	}
	//fmt.Println(employee)
	return employee
}

func EmployeeMethods(w http.ResponseWriter, r *http.Request) {
	/*
		switch r.Method {
		case "GET":
	*/
	fmt.Println("Test")
	var employee []dbModels.Employee
	employee = GetEmployees()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
	/*
		default:
			w.WriteHeader(http.StatusPreconditionFailed)
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	*/
}
