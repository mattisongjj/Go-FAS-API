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
