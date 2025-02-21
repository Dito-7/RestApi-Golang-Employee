package repository

import (
	"RestApi-Golang-Employee/model"
	"context"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepo struct {
	MongoCollection *mongo.Collection
}

func (r *EmployeeRepo) InserEmployee(emp *model.Employee) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), emp)

	if err != nil {
		slog.Info("Error while inserting employee")
		return nil, err
	}

	return result.InsertedID, nil
}

func (r *EmployeeRepo) FindEmployeeById(empID string) (*model.Employee, error) {
	var emp model.Employee

	err := r.MongoCollection.FindOne(context.Background(),
		bson.D{{Key: "employee_id", Value: empID}}).Decode(&emp)

	if err != nil {
		slog.Info("Error while finding employee by id")
		return nil, err
	}

	return &emp, nil
}

func (r *EmployeeRepo) FindAllEmployees() ([]model.Employee, error) {
	result, err := r.MongoCollection.Find(context.Background(), bson.D{})

	if err != nil {
		slog.Info("Error while finding all employees")
		return nil, err
	}

	var emps []model.Employee
	err = result.All(context.Background(), &emps)

	if err != nil {
		slog.Info("Error while decoding all employees")
		return nil, err
	}

	return emps, nil
}

func (r *EmployeeRepo) UpdateEmployee(empID string, updateEmp *model.Employee) (int64, error) {
	result, err := r.MongoCollection.UpdateOne(context.Background(),
		bson.D{{Key: "employee_id", Value: empID}},
		bson.D{{Key: "$set", Value: updateEmp}})

	if err != nil {
		slog.Info("Error while updating employee")
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (r *EmployeeRepo) DeleteEmployeeByID(empID string) (int64, error) {
	result, err := r.MongoCollection.DeleteOne(context.Background(),
		bson.D{{Key: "employee_id", Value: empID}})

	if err != nil {
		slog.Info("Error while deleting employee")
		return 0, err
	}

	return result.DeletedCount, nil
}

func (r *EmployeeRepo) DeleteAllEmployees() (int64, error) {
	result, err := r.MongoCollection.DeleteMany(context.Background(), bson.D{})

	if err != nil {
		slog.Info("Error while deleting all employees")
		return 0, err
	}

	return result.DeletedCount, nil
}
