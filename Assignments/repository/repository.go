package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"wifi/Assignments/models"
)

type EmployeeRepo struct {
	employees map[int]models.Employee
}

func NewEmployeeRepo() *EmployeeRepo {
	return &EmployeeRepo{employees: make(map[int]models.Employee)}
}

func (r *EmployeeRepo) CreateEmployee(w http.ResponseWriter, req *http.Request) {

	var emp models.Employee

	if err := json.NewDecoder(req.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}

	fmt.Println(emp)

	if _, exists := r.employees[emp.ID()]; exists {
		fmt.Println("Employee Already Exists")
		http.Error(w, "Invalid Request body", http.StatusUnprocessableEntity)
		return
	}

	r.employees[emp.ID()] = emp

	r.ListAllEmployees()

	w.WriteHeader(http.StatusCreated)

}

func (r *EmployeeRepo) GetEmployees(w http.ResponseWriter, req *http.Request) {

	idst := req.URL
	fmt.Printf("Idst %T idst %v", idst, idst)
	id := req.URL.Query().Get("id")

	employeeID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid Employee ID", http.StatusBadRequest)
		return
	}

	if emp, exists := r.employees[employeeID]; !exists {
		http.Error(w, "Employee Not Found", http.StatusNotFound)
	} else {

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(emp)
	}
}

func (r *EmployeeRepo) UpdateEmployee(w http.ResponseWriter, req *http.Request) {
	var emp models.Employee

	if err := json.NewDecoder(req.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}

	if _, exists := r.employees[emp.ID()]; !exists {
		http.Error(w, "Employee Not Found", http.StatusNotFound)
		return
	}

	r.employees[emp.ID()] = emp

	r.ListAllEmployees()

	w.WriteHeader(http.StatusNoContent)
}

func (r *EmployeeRepo) DeleteEmployee(w http.ResponseWriter, req *http.Request) {

	id := req.URL.Query().Get("id")

	employeeID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid Employee ID", http.StatusBadRequest)
		return
	}

	if _, exists := r.employees[employeeID]; !exists {
		http.Error(w, "Employee Not Found", http.StatusNotFound)
		return
	}

	delete(r.employees, employeeID)

	r.ListAllEmployees()

	w.WriteHeader(http.StatusNoContent)
}

func (r *EmployeeRepo) ListAllEmployees() {
	fmt.Println(r.employees)
}
