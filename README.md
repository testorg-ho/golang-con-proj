# Golang Console Project

This project is a console application written in Go that interacts with the OpsLevel CLI to retrieve a list of services and creates corresponding GitHub repositories using the GitHub CLI.

## Project Structure

```
golang-console-project
├── main.go              # Entry point of the application
├── internal
│   ├── opslevel
│   │   └── client.go    # OpsLevel client implementation
│   ├── github
│   │   └── client.go    # GitHub client implementation
│   └── services
│       └── manager.go   # Service manager logic
├── go.mod               # Module definition and dependencies
├── go.sum               # Module dependency checksums
└── README.md            # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd golang-console-project
   ```

2. **Install dependencies:**
   Ensure you have Go installed on your machine. Run the following command to download the necessary dependencies:
   ```
   go mod tidy
   ```

3. **Install OpsLevel CLI:**
   Follow the instructions on the [OpsLevel documentation](https://opslevel.com/docs) to install the OpsLevel CLI.

4. **Install GitHub CLI:**
   Follow the instructions on the [GitHub CLI documentation](https://cli.github.com/) to install the GitHub CLI.

## Usage

To run the application, execute the following command:
```
go run main.go
```

### Run Unit Tests

To execute unit tests, run:
```
go test ./... -v
```

### Run Integration Tests

To execute integration tests, run:
```
go test -tags=integration ./... -v
```

### Build the Project

To build the project, run:
```
go build -o golang-console-project
```

This will create an executable named `golang-console-project` in the current directory.

## Mock Generation

To generate mocks for the interfaces used in this project, use the following commands:

```bash
# Generate mock for OpsLevelClient
mockgen -source=internal/opslevel/client.go -destination=internal/opslevel/mocks/mock_client.go -package=mocks

# Generate mock for GitHubClient
mockgen -source=internal/github/client.go -destination=internal/github/mocks/mock_client.go -package=mocks
```

Ensure that `mockgen` is installed and available in your `PATH`. You can install it using:

```bash
go install github.com/golang/mock/mockgen@latest
```

Add the `GOPATH/bin` directory to your `PATH` if necessary:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.