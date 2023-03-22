package main

import (
	"database/sql"
	"fmt"
	"golang-webservice/routers"

	_ "github.com/lib/pq"
)

var PORT = "127.0.0.1:8080"

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "db-go-sql"
)

var (
	db  *sql.DB
	err error
)

type Employee struct {
	ID        int
	Full_Name string
	Email     string
	Age       int
	Division  string
}

func main() {
	// http.HandleFunc("/employees", getEmployees)
	// http.HandleFunc("/employee", createEmployees)
	// fmt.Println("Application is listening on port", PORT)
	// http.ListenAndServe(PORT, nil)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to database")
	// CreateEmployee()
	// GetEmployees()
	// UpdateEmployee()
	// DeleteEmployee()

	routers.StartServer().Run(PORT)
}

func CreateEmployee() {
	var employee = Employee{}
	sqlStatement := `
		INSERT INTO employees (id, full_name, email, age, division)
		VALUES($1, $2, $3, $4, $5)
		Returning *
	`
	err = db.QueryRow(sqlStatement, 4, "Airell Jordan", "airell@gmail.com", 23, "IT").
		Scan(&employee.ID, &employee.Full_Name, &employee.Email, &employee.Age, &employee.Division)
	if err != nil {
		panic(err)
	}
	fmt.Printf("New Employee Data : %+v\n", employee)
}

func GetEmployees() {
	var results = []Employee{}
	sqlStatement := `SELECT * FROM employees`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var employee = Employee{}
		err = rows.Scan(&employee.ID, &employee.Full_Name, &employee.Email, &employee.Age, &employee.Division)
		if err != nil {
			panic(err)
		}
		results = append(results, employee)
	}
	fmt.Println("Employee Data :", results)
}

func UpdateEmployee() {
	sqlStatement := `
		UPDATE employees
		SET full_name = $2, email = $3, age = $4, division = $5
		WHERE id = $1;
	`
	res, err := db.Exec(sqlStatement, 4, "Airell Jordan Wijaya", "airelljw@gmail.com", 21, "Data")
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Updated Data Amount :", count)
}

func DeleteEmployee() {
	sqlStatement := `
		DELETE FROM employees
		WHERE id = $1;
	`
	res, err := db.Exec(sqlStatement, 4)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted Data Amount :", count)
}

// var employees = []Employee{
// 	{ID: 1, Name: "Airell", Age: 23, Division: "IT"},
// 	{ID: 1, Name: "Nanda", Age: 27, Division: "Finance"},
// 	{ID: 1, Name: "Mailo", Age: 20, Division: "IT"},
// }

// func getEmployees(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	if r.Method == "GET" {
// 		json.NewEncoder(w).Encode(employees)
// 		return
// 	}
// 	http.Error(w, "Invalid method", http.StatusBadRequest)
// }

// func createEmployees(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	if r.Method == "POST" {
// 		name := r.FormValue("name")
// 		age := r.FormValue("age")
// 		division := r.FormValue("division")

// 		convertAge, err := strconv.Atoi(age)

// 		if err != nil {
// 			http.Error(w, "Invalid Age", http.StatusBadRequest)
// 			return
// 		}

// 		newEmployee := Employee{
// 			ID:       len(employees) + 1,
// 			Name:     name,
// 			Age:      convertAge,
// 			Division: division,
// 		}

// 		employees = append(employees, newEmployee)
// 		json.NewEncoder(w).Encode(newEmployee)
// 		return
// 	}

// 	http.Error(w, "Invalid method", http.StatusBadRequest)
// }
