# CRM API

This is a sample Customer Relationship Management (CRM) API built with Go. It provides basic functionalities to manage customer data, including creating, reading, updating, and deleting customer records. This API is designed to be a simple example that can be used for containerized application deployments or other educational purposes.

## Features

- **Get All Customers**: Retrieve a list of all customers.
- **Create Customer**: Add a new customer to the database.
- **Remove Customer**: Delete a customer from the database by ID.
- **Update Customer**: Update customer by ID.

## Endpoints

### Get All Customers

- **URL**: `/customers`
- **Method**: `GET`
- **Response**: JSON array of customer objects.

### Get A Customers
- **URL**: `/userid/{id}/`
- **Method**: `GET`
- **Response**: JSON array of customer objects.


### Create Customer
- **URL**: `/customer`
- **Method**: `POST`
- **Request Body**: JSON object representing the customer to be added.
- **Response**: JSON object of the created customer.

### Update Customer
- **URL**: `/customerupdate/{id}/`
- **Method**: `PUT`
- **Request Body**: JSON object representing the customer to be updated.
- **Response**: JSON object of customers list.

### Remove Customer
- **URL**: `/delete/customers/{id}`
- **Method**: `DELETE`
- **Response**: No content.

## Running the Application

### Prerequisites

- Go 1.22 or higher
- Docker (optional, for containerized deployment)

### Local Development

1. Clone the repository:
    ```sh
    git clone https://github.com/coles243/crm.git
    cd crm
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Run the application:
    ```sh
    go run main.go
    ```

4. The API will be available at `http://localhost:8080`.

### Containerized Deployment

1. Build the Docker image:
    ```sh
    docker build -t crm-api .
    ```

2. Run the Docker container:
    ```sh
    docker run -p -d 3000:3000 crm-api
    ```

3. The API will be available at `http://localhost:3000`.

## Sample Data

The application uses a JSON file (`database.json`) to store customer data. Here is an example of the JSON structure:

```json
[
  {
    "id": 1,
    "name": "Albert",
    "role": "Infrastructure Engineer",
    "email": "colealbert88@hotmail.com",
    "phone": "888-888-0000",
    "contacted": true
  },
  {
    "id": 2,
    "name": "Makri",
    "role": "Data Analyst",
    "email": "makrichevi1@aol.com",
    "phone": "871-888-2233",
    "contacted": true
  }
]