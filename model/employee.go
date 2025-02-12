package model

type Employee struct {
	EmployeeID string `bson:"employee_id" json:"employee_id,omitempty"`
	Name       string `bson:"name" json:"name,omitempty"`
	Department string `bson:"department" json:"department,omitempty"`
}
