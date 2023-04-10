# Helper App Backend

The Helper App Backend is a Golang-based RESTful API that serves as the backend support for the Helper App. The API is built with the `Gin` web framework and uses `MongoDB` for data storage.

## Requirements

To run this project, you need to have the following software installed on your machine:

- `Go 1.18` or later
- `MongoDB`

## Installation

To install the Helper App Backend, you can clone this repository and then build and run the application. Here are the steps:

1. Clone the repository:
`git clone https://github.com/supperdoggy/helper-backend.git`

2. Change to the project directory:
`cd helper-backend`

3. Install the required packages:
`go mod download`

4. Build and run the application:
`go run main.go`


The server should now be running on `http://localhost:8080`.

## API Endpoints

The Helper App Backend provides several API endpoints for managing users, authentication, and adverts. Please refer to the Postman collection available in this repository for detailed information on how to call these endpoints.

### Postman Collection

A Postman collection with all the Helper App Backend API endpoints and examples of their calls is available in this repository. Please import the collection into your Postman application to explore the API functionalities.

The collection can be found at `./helper-backend.postman_collection.json`.

