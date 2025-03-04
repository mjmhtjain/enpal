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
The server will start on `localhost:8080` by default.

### Using Docker
1. Build the Docker image:
```bash
make docker-build
```

2. Run the application in Docker:
```bash
make docker-run
```

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

