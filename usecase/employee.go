package usecase

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeUsecase struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (svc *EmployeeUsecase) CreateEmployeeHandler(w http.ResponseWriter, r *http.Request) {}
func (svc *EmployeeUsecase) GetEmployeeID(w http.ResponseWriter, r *http.Request)         {}
func (svc *EmployeeUsecase) GetAllEmployee(w http.ResponseWriter, r *http.Request)        {}
func (svc *EmployeeUsecase) UpdateEmployeeID(w http.ResponseWriter, r *http.Request)      {}
func (svc *EmployeeUsecase) DeleteEmployeeByID(w http.ResponseWriter, r *http.Request)    {}
func (svc *EmployeeUsecase) DeleteAllEmployee(w http.ResponseWriter, r *http.Request)     {}
