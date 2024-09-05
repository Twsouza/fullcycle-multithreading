
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

Run the program:

```bash
go run main.go <cep>
```

Example:

```bash
go run main.go 01153000
```

## Output

The command-line output will display:
- Address details (from the fastest API response).
- The name of the API that provided the response.
- A timeout error message if no response is received within the specified limit.
