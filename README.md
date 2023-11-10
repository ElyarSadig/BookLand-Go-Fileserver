# BookLand File Server in Go
A file server designed for the BookLand project in Go, providing efficient file management capabilities.
## Overview
This file server facilitates the storage and management of various file types associated with the BookLand project. It offers functionalities for uploading and accessing files categorized into distinct directories.
## Features
- **Directory Structure**: Organized directory structure for different file types, such as identities, publications, book covers, and books.
- **Upload Handlers**: Dedicated upload handlers for each file type, ensuring secure and controlled file uploads.
- **File Type Verification**: Ensures the uploaded files adhere to supported file types (PDF, PNG, JPG, JPEG).
- **Access Control**: Only trusted domains can access files. Uploading is restricted to authenticated backends using auth tokens.
- **Server Configuration**: Runs on port 8080 by default.
## Getting Started
### Prerequisites
Ensure you have Go installed on your system.
### Installation
1. Clone the repository
2. Run `go build` to build the project.
3. Execute the generated binary to start the file server.
### Docker Integration
For containerized deployment using Docker, use the provided Dockerfile and Docker Compose file:

**Dockerfile**

The Dockerfile included in the repository facilitates the containerization of the BookLand file server. Build the image and run the container using the following commands:
  ```bash
# Build the Docker image
docker build -t bookland-file-server .
  
# Run the Docker container
docker run -p 8080:8080 bookland-file-server
```

**Docker Compose**

The docker-compose.yml file simplifies orchestration and deployment. Run the BookLand file server with Docker Compose using:

```
docker-compose up
```

## Usage

**Uploading Files**
- To upload files, only authenticated backends with an auth token can use the designated endpoints for each file type. For example:
  - **/upload/identities** for identity files.
  - **/upload/publications** for publication files.
  - **/upload/book-covers** for book cover images.
  - **/upload/books** for book files.

**Accessing Files**
- Files can be accessed using the respective endpoints based on file types, accessible only by trusted domains:
  - **/identities/** for identity files.
  - **/publications/** for publication files.
  - **/book-covers/** for book cover images.
  - **/books/** for book files.

## Configuration
Adjust server configurations, trusted domain lists, or authentication mechanisms in the main.go and middleware.go file as needed.

## Contributing
Feel free to fork this repository and submit pull requests for any enhancements or bug fixes.
