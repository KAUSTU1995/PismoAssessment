# Project Documentation

## Overview

This project is designed to provide a [brief description of the project]. It utilizes Docker for containerization and the Go programming language for the application logic. This document outlines how to build, run, and test the application, as well as how to generate Swagger documentation.

## Prerequisites

Before you start, ensure you have the following installed:

- **Docker**: [Download and install Docker](https://www.docker.com/get-started)
- **Docker Compose**: Included with Docker Desktop
- **Go**: [Download and install Go](https://golang.org/doc/install)
- **Make**: The `make` command is typically included in build-essential packages. Install it via:

    - **Debian/Ubuntu**:

      ```bash
      sudo apt-get install build-essential
      ```

    - **Fedora**:

      ```bash
      sudo dnf install make
      ```

    - **MacOS**:

      ```bash
      brew install make
      ```


## Project Structure

- `cmd/`: Contains the main application entry points.
- `controllers/`: Contains the application's controllers.
- `models/`: Contains data models.
- `errors/`: Contains custom error definitions.
- `docs/`: Output directory for Swagger documentation.

## Makefile Commands

This project uses a Makefile to manage the application's lifecycle. The following commands are available:

### 1. Build Docker Images

Build the Docker images specified in the `docker-compose.yml` file.

    make build

### 2. Run Docker Containers

Start the Docker containers in detached mode.

    make run

### 3. Stop Docker Containers

Stop and remove the running Docker containers.

    make stop

### 4. Run Tests

Execute all tests in the project with verbose output.

    make test

### 5. Generate Swagger Documentation

Generate Swagger documentation based on the application's code and specified directories.

    make swag

This command will create or update the Swagger documentation in the `docs` directory.

## How to Use

1. **Clone the Repository**: First, clone the repository to your local machine.

   git clone <repository-url>
   cd <project-directory>

2. **Build the Docker Images**: Run the build command to create the Docker images.

   make build

3. **Run the Application**: Start the application by running the Docker containers.

   make run

4. **Access the Application**: Navigate to `http://localhost:<port>` in your web browser to access the application. Typically 8080.

5. **Run Tests**: You can run tests at any time using the following command:

   make test

6. **Generate Documentation and Access Swagger**: To create or update the Swagger documentation, run:

   make swag

   After generating the documentation, access it via the URL: `http://localhost:8080/swagger/index.html`.

## Additional Information

- Ensure your Docker daemon is running before executing the commands in this document.
- Modify the `docker-compose.yml` file as needed to suit your application's requirements.
- If you need to change configurations, you can modify the `config.json` or other relevant configuration files in the project.
