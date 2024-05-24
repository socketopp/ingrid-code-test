# Coding Assignment

## Introduction

This repository contains the solution for a coding assigment sent out by Ingrid. The assignment involves creating a REST API in Go to calculate the duration and distance between multiple destinations and a source point using the OSRM API.

### Project structure
```
coding-assignment
├─ Dockerfile
├─ go.mod
├─ main.go
├─ osrm
│  └─ request.go
├─ README.md
└─ server
   ├─ handlers.go
   └─ mux.go
```

## Getting Started

### Prerequisites

- Go 1.22.3 or higher
- Docker (optional, for containerized deployment)

### Installation

1. Clone the repository to your local machine:

   ```bash
   $ git clone https://github.com/socketopp/coding-assignment.git
   ```
2. Navigate to the project directory:
    ```bash 
    $ cd coding-assignment`
    ```
    
### Running Locally

To run the application locally, follow these steps:

1. Build the Go executable:
    ```bash
    $ go build -o main .
    ```
2. Start the server:
   ```bash
    $ ./main
    ```
    
    Alternatively just run the main file without building
    ```bash
    $ go run main.go
    ```
    
### Running with Docker
##### Prerequisites
Make sure the docker daemon is running.

To run the application using Docker, follow these steps:

1. Build the Docker image:

    ```bash
    $ docker build -t ingrid-go-image .
    ```
2. Run the Image in a Docker container:
    
    ```bash
    $ docker run -d --rm --name ingrid-go-container -p 8080:8080 ingrid-go-image
    ```
3. Check the logs
     ```bash
    $ docker logs -f ingrid-go-container
    2024/05/23 08:02:09 Starting server on :8080
    Building REST API's
     ```

### Deploy to Google Cloud Run
Ensure you have Google Cloud SDK installed and authenticated and pushed your docker image to GCP container registry before, then proceed with:

  1. Tag and push your Docker image to Google Container Registry:
      ```bash
      $ docker tag ingrid-go-image gcr.io/your-project-id/ingrid-go-image:v1
      $ docker push gcr.io/your-project-id/ingrid-go-image:v1
      ```

  2. Deploy the image to Cloud Run:
      ```bash
      $ gcloud run deploy ingrid-go-service --image gcr.io/your-project-id/ingrid-go-image:v1 --platform managed
      ```
        
### Testing the API

You can test the API using any HTTP client or tool like cURL. Here's an example cURL command to retrieve route information:
```bash
curl 'http://localhost:8080/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219'
{"source":"13.388860,52.517037","routes":[{"destination":"13.397634,52.529407","distance":1886.8,"duration":260.3},{"destination":"13.428555,52.523219","distance":3804.2,"duration":389.3}]}
```

### API Documentation
**Endpoint: /routes**

    Method: GET
    Parameters:
        src: Source point coordinates (latitude,longitude)
        dst: Array of destination points coordinates (latitude,longitude)

Example Request:

```http

GET /routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219 HTTP/1.1
Host: localhost:8080
```
Example Response:

```json
{
  "source": "13.388860,52.517037",
  "routes": [
    {
      "destination": "13.397634,52.529407",
      "distance": 1886.8,
      "duration": 260.3
    },
    {
      "destination": "13.428555,52.523219",
      "distance": 2221.6,
      "duration": 322.5
    }
  ]
}
```
