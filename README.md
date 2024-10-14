
# Multithreaded API Request - Fastest Response in Go

This project is implemented in Go and demonstrates how to make concurrent API requests using multithreading, selecting the fastest response between two APIs that return address data based on a provided zip code (CEP).

### API Details:

- Brasil API
- ViaCEP

### Project Overview:

- **Multithreading**: The project makes simultaneous requests to both APIs using Go routines and waits for the first response.
- **Response Handling**: The fastest response is accepted and displayed in the command line, including the API that provided the result.
- **Timeout**: If neither API responds within 1 second, a timeout error is triggered, and the error is displayed in the command line.

## How to Run

To run this project, you need to have Go installed on your machine.

### Using `go run`

Run the program:

```bash
go run *.go <cep>
```

Replace `<cep>` with the zip code you want to search for.

```bash
go run *.go 01311-000
```

### Using Makefile

Alternatively, you can use the Makefile to run the program:

```bash
make run cep=01311-000
```

Replace `01311-000` with the zip code you want to search for.

```bash
make run cep=01311-000
```

To clean the build artifacts, run:

```bash
make clean
```

This section provides instructions for both running the program directly with `go run` and using the Makefile to build and run the program.

## Output

The command-line output will display:
- Address details (from the fastest API response).
- The name of the API that provided the response.
- A timeout error message if no response is received within the specified limit.
