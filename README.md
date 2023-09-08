# Api.Crud.Go.Gin

This is a Go project that implements a CRUD (Create, Read, Update, Delete) for managing student information. It uses the Gin framework to create HTTP endpoints and Gorm to interact with a PostgreSQL database. Below are the project details:

## Endpoints

- **Health Check**: Checks the status of the application.
  - **Method**: GET
  - **Endpoint**: `/`
  - **Controller**: `controllers.HealthCheck`

- **List All Students**: Returns a list of all registered students.
  - **Method**: GET
  - **Endpoint**: `/students`
  - **Controller**: `controllers.ListAll`

- **Create a New Student**: Creates a new student record.
  - **Method**: POST
  - **Endpoint**: `/students`
  - **Controller**: `controllers.CreateNewStudent`

- **Search Student by ID**: Retrieves information about a student based on their ID.
  - **Method**: GET
  - **Endpoint**: `/students/:id`
  - **Controller**: `controllers.SearchStudentById`

- **Delete Student by ID**: Deletes a student based on their ID.
  - **Method**: DELETE
  - **Endpoint**: `/students/:id`
  - **Controller**: `controllers.DeleteStudent`

- **Update Student by ID**: Updates information about a student based on their ID.
  - **Method**: PATCH
  - **Endpoint**: `/students/:id`
  - **Controller**: `controllers.UpdateStudent`

- **Search Student by CPF**: Retrieves information about a student based on their CPF.
  - **Method**: GET
  - **Endpoint**: `/students/cpf/:cpf`
  - **Controller**: `controllers.SearchStudentByCPF`

## Data Model

The data model used in the project is defined as follows:

```go
package models

import "gorm.io/gorm"

type Student struct {
    gorm.Model
    Name string `json:"name"`
    CPF  string `json:"cpf"`
    Mail string `json:"mail"`
}
```

## Database Configuration

This project uses a PostgreSQL database. Make sure to configure the database connection information in the corresponding configuration file.

## How to Run the Project

1. Clone the repository to your local machine.
2. Ensure you have Go installed on your machine.
3. Install project dependencies using `go get`.
4. Configure the database connection information in the configuration file.
5. Run the project with the `go run main.go` command.

Make sure all dependencies are properly installed before running the project.

## Contributions

Contributions are welcome! Feel free to open issues and send pull requests to improve this project.

Thank you for using this project!