# Customer CRUD Application

This is a GoLang application that provides CRUD (Create, Read, Update) operations for managing customer data. It includes a web interface with three pages: index.html, create.html, and edit.html. The index.html page displays all customers and supports search and sorting operations.

## Prerequisites

Before running the application, ensure that you have the following dependencies installed:

- Go programming language
- Docker

## Setup

1. Clone the repository to your local machine:

```shell
https://github.com/duyguseyhan/CrudOperationsInGo.git
```

2. Change into the project directory.


3. Start the PostgreSQL database using Docker Compose:

```shell
docker-compose up -d
```

This command will launch a PostgreSQL container with the required configuration.

4. Build and run the application:

If you are using **Windows**:

```shell
go build
go run ./
```

If you are using **macOS**:

```shell
go build
go run *.go
```

The application will compile and start running. You can access it in your web browser at http://localhost:8080.

## Usage

- Open your web browser and navigate to http://localhost:8080.
- The index.html page will display a list of all customers.
- Use the search and sorting features to filter and order the customer data.
- To create a new customer, click the "Create" button, which will take you to create.html.
- To edit an existing customer, click the "Edit" button next to the desired customer on the index.html page. You will be redirected to edit.html.
- On the create.html and edit.html pages, fill in the necessary details and click the related button to save the changes.

## Troubleshooting

- If you encounter any issues with the application, please make sure that the PostgreSQL container is running and accessible.
- Check that the necessary Go packages are installed correctly.
- Verify that the database connection settings in the application code are correct.

---
