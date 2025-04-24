# go-exam-p1

## Overview
This project demonstrates basic Go programming concepts and the use of the Gin framework to create a simple REST API. It is designed for beginners who are learning Go and want to understand how to build and validate RESTful APIs.

## Features
- A `User` struct to represent user data with fields for name, email, and age.
- Functions to update user email and validate user data.
- A REST API endpoint `/user` implemented using the Gin framework:
  - Accepts JSON input with `name` and `email` fields.
  - Validates the input data.
  - Returns the created user with a `201 Created` status.

## Prerequisites
- Go installed on your system (version 1.16 or later).
- Basic understanding of Go programming.

## Installation
1. Clone this repository:
   ```bash
   git clone https://github.com/witchakornb/go-exam-p1.git
   ```
2. Navigate to the project directory:
   ```bash
   cd go-exam-p1
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage
1. Run the application:
   ```bash
   go run main.go
   ```
2. The server will start on `http://localhost:8080`.

3. Test the `/user` endpoint:
   - Use a tool like `curl` or Postman to send a POST request to `http://localhost:8080/user`.
   - Example request:
     ```bash
     curl -X POST http://localhost:8080/user \
     -H "Content-Type: application/json" \
     -d '{"name": "John Doe", "email": "john.doe@example.com"}'
     ```
   - Example response:
     ```json
     {
       "name": "John Doe",
       "email": "john.doe@example.com"
     }
     ```

## Code Explanation
### main.go
- **User Struct**: Represents user data with fields for `name`, `email`, and `age`.
- **updateUserEmail Function**: Updates the email field of a `User` struct.
- **validateUser Function**: Validates the `User` struct to ensure the email is not empty and the age is a positive integer.
- **Gin Router**: Defines a POST `/user` endpoint to handle user creation.

## Project Structure
```
.
├── go.mod       # Module dependencies
├── go.sum       # Dependency checksums
├── LICENSE      # License file
├── main.go      # Main application code
├── README.md    # Project documentation
```

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request.

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.

## Acknowledgments
- [Gin Web Framework](https://github.com/gin-gonic/gin) for making REST API development in Go simple and efficient.