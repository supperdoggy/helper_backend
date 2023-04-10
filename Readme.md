Helper App Backend

The Helper App Backend is a Golang-based RESTful API that serves as the backend support for the Helper App. The API is built with the Gin web framework and uses MongoDB for data storage. The main purpose of the API is to allow users to create adverts about help that they need or can provide to others.

Requirements

To run this project, you need to have the following software installed on your machine:

Go 1.18 or later
MongoDB
Installation

To install the Helper App Backend, you can clone this repository and then build and run the application. Here are the steps:

Clone the repository:
bash
Copy code
git clone https://github.com/supperdoggy/helper-backend.git
Change to the project directory:
bash
Copy code
cd helper-backend
Install the required packages:
go
Copy code
go mod download
Build and run the application:
go
Copy code
go run main.go
The server should now be running on http://localhost:8080.

API Endpoints

The Helper App Backend provides several API endpoints for managing users, authentication, and adverts. You can find the examples of all endpoint calls in the Postman collection available in the postman/ directory.

Postman Collection

The Postman collection for the Helper App Backend contains examples of all the API endpoints. To use it, you need to have Postman installed on your machine.

Open Postman and import the collection by clicking on File -> Import and selecting the helper-backend.postman_collection.json file in the postman/ directory.
You can now use the imported collection to test the API endpoints by selecting the desired endpoint and clicking on Send. The responses from the server will be displayed in the Response tab.
Conclusion

The Helper App Backend is a simple yet powerful RESTful API that provides the backend support for the Helper App. With its easy-to-use endpoints and efficient data storage, it can help users create adverts about help they need or can provide to others.
