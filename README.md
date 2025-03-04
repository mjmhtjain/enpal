# Enpal Application

## Prerequisites
- Go installed on your system
- Docker and Docker Compose (for running with Docker)

## Setup
1. Clone the repository
2. Install dependencies:
```bash
go mod download
```

## Running the Application

### Local Development
1. Start the application locally:
```bash
make dev
```

### Using Docker
1. Build the Docker image:
```bash
make docker-build
```

2. Run the application in Docker:
```bash
make docker-run
```
The server will start on `localhost:3000` by default.

3. Stop the Docker containers:
```bash
make docker-stop
```

Additional Docker commands:
- Rebuild Docker image: `make docker-rebuild`
- Clean up Docker resources: `make docker-cleanup`
- Remove unused Docker containers: `make docker-prune`

## Development Commands
- Format code: `make fmt`
- Run linter: `make lint`
- Run code vetting: `make vet`
- Clean build artifacts: `make clean`
- To generate mocks for testing, run:
```bash
mockery --name=IAppointmentRepo \
--dir=src/internal/repository \
--output=src/internal/mocks \
--structname=AppointmentRepo
```

## Testing
Run tests with:
```bash
make test
```

Generate test coverage report:
```bash
make test-coverage
```

## Available Make Commands
Run `make help` to see all available commands and their descriptions.

## API Endpoints

### Health Check
- **GET** `/health`
- Returns the health status of the application
- Response:
  ```json
  {
    "status": "healthy"
  }
  ```
- Status Codes:
  - `200`: Application is healthy

### Calendar Query
- **POST** `/calendar/query`
- Returns available appointment slots based on the provided date, products, language, and rating
- Request Body:
  ```json
  {
    "date": "2024-05-03",
    "products": ["Heatpumps"],
    "language": "English",
    "rating": "Silver"
  }
  ```
- Response: Array of available slots with count and start time
  ```json
  [
    {
      "available_count": 1,
      "start_date": "2024-05-03T10:30:00.000Z"
    },
    {
      "available_count": 1,
      "start_date": "2024-05-03T11:00:00.000Z"
    }
  ]
  ```
- Status Codes:
  - `200`: Successfully retrieved available slots
  - `400`: Bad request - Invalid input parameters
  - `500`: Internal server error


