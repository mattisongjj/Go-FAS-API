# Go-FAS-API

## Project Overview
Go-FAS-API is a backend service written in Go that manages financial assistance schemes for needy individuals and families. This service provides RESTful API endpoints for managing applicants, schemes, and applications.

## Prerequisites
Ensure you have met the following requirements:
- Go (version 1.16 or higher)

## Installation and running the application

1. **Clone the Repository**
   ```bash
   git clone git@github.com:mattisongjj/Go-FAS-API.git
   cd Go-FAS-API
2. **Build and run the application**
   ```bash
    go run cmd/api/main.go
  The application will start and listen on http://localhost:8080
 
## Authentication

As a placeholder for authentication in a large scale app, to authenticate a request via a middleware function, follow the steps below:

1. **Send the Username**: Include the username "john" in the URL parameters.
   - Example: `http://localhost:8080/api/applicants?username=john`

2. **Include Authorization Header**: Add an `Authorization` header with the value `123456ABC` in the request.
   - Example:
     ```
     "Authorization": "123456ABC"
     ```
## API Documentation
The API documentation for this application can be found [here](https://documenter.getpostman.com/view/33827285/2sAXjRW9NG)

### Endpoints Overview

- **Applicants**
  - `GET /api/applicants`: Get all applicants.
  - `POST /api/applicants`: Create a new applicant.
  
- **Schemes**
  - `GET /api/schemes`: Get all schemes.
  - `POST /api/schemes`: Create a new scheme.
  - `GET /api/schemes/eligible?applicant={id}`: Get eligible schemes for a specific applicant.

- **Applications**
  - `GET /api/applications`: Get all applications.
  - `POST /api/applications`: Create a new application.

### Making POST Request

When making POST requests, ensure that the request body is formatted in JSON as per the following structures:

#### Applicant

```json
{
  "id": "string",
  "name": "string",
  "employment_status": "employed | unemployed",
  "sex": "male | female",
  "date_of_birth": "YYYY-MM-DD",
  "household": [
    {
      "id": "string",
      "applicant_id": "string",
      "name": "string",
      "employment_status": "employed | unemployed",
      "sex": "male | female",
      "date_of_birth": "YYYY-MM-DD",
      "relation": "mother | father | spouse | sister | brother | daughter | son | other"
    }
  ]
}
```
#### Scheme

```json
{
  "id": "string",
  "name": "string",
  "criteria": {
    "employment_status": "employed | unemployed",
    "has_children": "true | false"
  },
  "benefits": [
    {
      "id": "string",
      "scheme_id": "string",
      "name": "string",
      "amount": "number"
    }
  ]
}
```
#### Application
```json
{
  "id": "string",
  "applicant_id": "string",
  "scheme_id": "string",
  "status": "string"
}
```


