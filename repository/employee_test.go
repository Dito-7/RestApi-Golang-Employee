package repository

import (
	"RestApi-Golang-Employee/model"
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.ytntn.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))

	if err != nil {
		log.Fatal("Error while connecting to mongo", err)
	}

	log.Println("Connected to MongoDB")

	err = mongoTestClient.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error while pinging to mongo", err)
	}

	log.Println("Ping to MongoDB successful")

	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	//dummy data
	emp1 := uuid.New().String()
	emp2 := uuid.New().String()

	// Connect to the collection
	coll := mongoTestClient.Database("companydb").Collection("employee_test")

	empRepo := EmployeeRepo{MongoCollection: coll}

	// Insert employee
	t.Run("InsertEmployee 1", func(t *testing.T) {
		emp := model.Employee{
			Name:       "John",
			Department: "IT",
			EmployeeID: emp1,
		}

		result, err := empRepo.InserEmployee(&emp)

		if err != nil {
			t.Fatal("Error while inserting employee", err)
		}

		t.Log("Inserted employee with ID: ", result)
	})
	t.Run("InsertEmployee 2", func(t *testing.T) {
		emp := model.Employee{
			Name:       "John",
			Department: "IT",
			EmployeeID: emp2,
		}

		result, err := empRepo.InserEmployee(&emp)

		if err != nil {
			t.Fatal("Error while inserting employee", err)
		}

		t.Log("Inserted employee with ID: ", result)
	})

	//get employee
	t.Run("FindEmployeeById", func(t *testing.T) {
		result, err := empRepo.FindEmployeeById(emp1)

		if err != nil {
			t.Fatal("Error while finding employee by id", err)
		}

		t.Log("Found employee: ", result)
	})

	//get all employees
	t.Run("FindAllEmployees", func(t *testing.T) {
		result, err := empRepo.FindAllEmployees()

		if err != nil {
			t.Fatal("Error while finding all employees", err)
		}

		t.Log("Found all employees: ", result)
	})

	//update employee
	t.Run("UpdateEmployee", func(t *testing.T) {
		emp := model.Employee{
			Name:       "John Doe Aka John",
			Department: "Business",
			EmployeeID: emp1,
		}

		result, err := empRepo.UpdateEmployee(emp1, &emp)

		if err != nil {
			t.Fatal("Error while updating employee", err)
		}

		t.Log("Updated employee with ID: ", result)
	})

	// get employee after update
	t.Run("Get Employee After Update", func(t *testing.T) {
		result, err := empRepo.FindEmployeeById(emp1)

		if err != nil {
			t.Fatal("Error while finding employee by id", err)
		}

		t.Log("Found employee: ", result)
	})

	//delete employee
	t.Run("DeleteEmployeeByID", func(t *testing.T) {
		result, err := empRepo.DeleteEmployeeByID(emp1)

		if err != nil {
			t.Fatal("Error while deleting employee", err)
		}

		t.Log("Deleted employee with ID: ", result)
	})

	//get employee after delete
	t.Run("Get Employee After Delete", func(t *testing.T) {
		result, err := empRepo.FindEmployeeById(emp1)

		if err == nil {
			t.Fatal("Employee still exists after deletion", err)
		}

		t.Log("Employee", result)
	})

	//delete all employees
	t.Run("DeleteAllEmployees", func(t *testing.T) {
		result, err := empRepo.DeleteAllEmployees()

		if err != nil {
			t.Fatal("Error while deleting all employees", err)
		}

		t.Log("Deleted all employees: ", result)
	})
}
