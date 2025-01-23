package service

import (
	"net/http"
)

type EmployeeService interface {
	CreateEmployee(w http.ResponseWriter, r *http.Request) //Create
	GetEmployees(w http.ResponseWriter, r *http.Request)   //read
	UpdateEmployee(w http.ResponseWriter, r *http.Request) //update
	DeleteEmployee(w http.ResponseWriter, r *http.Request) // delete
	ListAllEmployees()
}
