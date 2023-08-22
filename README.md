# TODO ✅

## Introduction

This project is a personal learning exercise in building a REST API using modern technologies and best practices. The aim of this project is to gain hands-on experience in designing, implementing, and testing a production-grade API. The API will provide basic functionality for managing tasks and lists, as well as user authentication. 🚀📚

In addition to the main functionality, the API will also handle static files such as images. This will provide a complete experience in building a REST API that includes serving static files. 🖼️

## Installation

To run this project, you will need the following tools:

### Docker 🐳

Docker is a tool used to automate the deployment of applications within containers. To install Docker, follow these steps:

1. Go to the [official website](https://www.docker.com/products/docker-desktop) and download the package for your operating system.
2. Follow the installation instructions provided in the package.

## Running the Project 🏃‍♂️

To run this project, follow these steps:

1. Clone the repository to your local machine.
2. Navigate to the project's root directory.
3. Create a `.env` file with the necessary environment variables (see the Environment Variables section above for details).
4. Run `docker compose build` to build the Docker images.
5. Run `docker compose up` to start the containers.

Once the containers are running, you can access the API via `http://localhost:8000`. 🚀🌐

## Environment Variables 🔧

The following environment variables must be declared to run the API:

- `PORT`: The port number on which the server will listen.
- `SRV_HOST`: The hostname for the server.
- `SRV_SECRET_KEY`: A secret key used for encrypting session data.
- `SRV_ALLOWED_ORIGINS`: The allowed host origins to use this server.
- `DB_USER`: The username for the database connection.
- `DB_PASSWORD`: The password for the database connection.
- `DB_NAME`: The name of the database.
- `DB_HOST`: The host of the database.
- `DB_PORT`: The port of the database.

Here is an example `.env` file:

```
PORT=8000
SRV_HOST=localhost
SRV_SECRET_KEY=mysecretkey
SRV_ALLOWED_ORIGINS=*
DB_USER=todo
DB_PASSWORD=mysecretpassword
DB_NAME=tasks_db
DB_HOST=localhost
DB_PORT=5432
```

## API Reference 📖

The API provides a range of useful endpoints for developers to leverage. In addition to the currently available endpoints, there are plans to expand the API with new endpoints in upcoming releases. Some of the current endpoints include those for user authentication, data retrieval, and data modification. As the API continues to evolve, it will be important to stay up to date on the latest developments to ensure that your application is making the most of the available endpoints. 🚀📚

### TASKS 📝

### GET /tasks

This endpoint returns all the tasks belonging to an account in the database.

### GET /tasks/{id}

Returns a specific task with the given {id}.

### POST /tasks

Creates a new task with the given parameters:

- task_name
- task_description
- task_due_date
- task_status

### PUT /tasks/{id}

Updates an existing task with the given {id}. The following parameters can be updated:

- title
- description
- list_id

### DELETE /tasks/{id}

Deletes a specific task with the given {id}.

### LISTS 📚

### GET /lists

Returns all the lists in the database

.

### GET /lists/{id}

Returns a specific list with the given {id}.

### POST /lists

Creates a new list with the given parameters:

- title
- description
- background_picture_url
- created_by

### PUT /lists/{id}

Updates an existing list with the given {id}. The following parameters can be updated:

- title
- description
- background_picture_url

### DELETE /lists/{id}

Deletes a specific list with the given {id}.

### AUTHENTICATION 🔒

### POST /register

Registers a new user with the given parameters:

- name
- last_name
- nickname
- email
- password

### POST /login

Authenticates a user with the given credentials:

- nickname
- email
- password

## Error Handling ❗🚦

This project implements error handling at every layer of the architecture, including the database layer, database implementation, server layer, and server implementation. This ensures that errors are caught and handled appropriately, improving the reliability and stability of the API. Specific error messages are returned to the client to help with debugging and troubleshooting.

## Conclusion 🎉✅

This API is very simple, but it has all the necessary functionality to perform CRUD operations on lists, tasks, and basic authentication services. It provides a great foundation for building more advanced features and expanding the capabilities of your application. Keep exploring and learning!
