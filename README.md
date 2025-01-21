# 🛠️ Go Categories Microservice

This is a simple microservice built with Go as a learning project. The primary goal is to explore Go's syntax, principles, and best practices while building a functional system for managing categories.

---

## 📖 About the Project

- **Purpose**:  
  To practice Go programming and understand its features like modular design, error handling, and clean architecture.

- **Features**:  
  - Basic implementation of category creation.  
  - Validation for category inputs.  
  - Logging system to track operations and errors.  

---

## 🚀 Technologies

- **Go**:  
  The primary language used to develop this microservice, providing simplicity and efficiency.

- **Clean Architecture**:  
  The code structure follows Clean Architecture principles to ensure modularity and maintainability, making it easier to extend or modify the microservice.

---

## 📂 Project Structure

The project is organized as follows:

```plaintext
go-categories-microservice/
├── cmd/
│   └── api/
│       ├── controllers/
│       │   ├── create-category.go
│       │   ├── list-categories.go
│       │   └── routes.go
│       ├── main.go
│       └── routes.go
├── internal/
│   ├── entities/
│   │   └── category.go
│   ├── repositories/
│   │   ├── in-memory-repository.go
│   │   └── interface.go
│   └── use-cases/
│       ├── create-category.go
│       └── list-categories.go
├── pkg/
├── tmp/
├── .air.toml
├── .gitignore
├── api.http
├── go.mod
├── go.sum
└── README.md
```

### Explanation of Folders and Files

- **`cmd/api/`**:  
  Contains the entry point of the application and routes configuration.  
  - `controllers/`: Contains controllers to handle HTTP requests and responses.  
  - `main.go`: The main entry point of the application.

- **`internal/`**:  
  Encapsulates the core business logic of the microservice.  
  - `entities/`: Defines core entities like `category.go`.  
  - `repositories/`: Contains data storage logic, including interfaces and implementations.  
  - `use-cases/`: Implements specific use cases like creating and listing categories.

- **`pkg/`**:  
  Contains shared utility functions or packages that are reusable across the project.

- **`tmp/`**:  
  Temporary files or testing-related artifacts.

- **`.air.toml`**:  
  Configuration file for hot reloading during development.

- **`api.http`**:  
  A file to test the API endpoints using HTTP requests.

---

## ⚙️ API Endpoints

### Health Check
- **URL**: `/healthy`  
- **Method**: `GET`  
- **Response**:
----
{
    "status": "ok"
}
----

### List Categories
- **URL**: `/categories`  
- **Method**: `GET`  
- **Response**:
```json
[
    {
        "id": 1,
        "name": "Category 1"
    },
    {
        "id": 2,
        "name": "Category 2"
    }
]
```

### Create a Category
- **URL**: `/categories`  
- **Method**: `POST`  
- **Request Body**:
```json
{
    "Name": "Golang"
}
```
- **Response**:
```json
{
    "id": 3,
    "name": "Golang"
}
```

---

## 📚 Examples

### Sample `category.go` Entity
```json
package entities

type Category struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
```

### Sample `create-category.go` Use Case
```json
package usecases

import (
    "errors"
    "go-categories-microservice/internal/entities"
    "go-categories-microservice/internal/repositories"
)

func CreateCategory(repo repositories.CategoryRepository, name string) (*entities.Category, error) {
    if name == "" {
        return nil, errors.New("name cannot be empty")
    }
    category := &entities.Category{Name: name}
    return repo.Save(category)
}
```

### Sample `routes.go`
```json
package controllers

import (
    "net/http"
    "github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/healthy", HealthCheck).Methods("GET")
    router.HandleFunc("/categories", ListCategories).Methods("GET")
    router.HandleFunc("/categories", CreateCategory).Methods("POST")
}
```

---

## 🔧 Running the Project

1. **Clone the Repository**:
```bash
git clone https://github.com/your-username/go-categories-microservice.git
cd go-categories-microservice
```

2. **Install Dependencies**:
Ensure you have Go installed and run:
```bash
go mod tidy
```

3. **Run the Application**:
```bash
go run cmd/api/main.go
```

4. **Test API Endpoints**:  
Use the provided `api.http` file to test endpoints. For example:
```json
GET http://localhost:8080/healthy

GET http://localhost:8080/categories

POST http://localhost:8080/categories
Content-Type: application/json

{
    "Name": "Golang"
}
```

---

## 🛡️ Key Principles

- **Validation**: Ensures input data for category creation meets specified rules.  
- **Error Handling**: Structured error management for better debugging and resilience.  
- **Logging**: Tracks application events, errors, and operations for monitoring and debugging.  

---

## 🔗 Resources

- [Go Documentation](https://go.dev/doc/)  
- [Clean Architecture Overview](https://github.com/jasontaylordev/CleanArchitecture)
