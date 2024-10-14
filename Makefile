# Define the name of the executable
EXECUTABLE := myprogram

# Define the source files
SOURCES := $(wildcard *.go)

# Build the executable
build:
    go build -o $(EXECUTABLE) $(SOURCES)

# Run the program
run: build
    ./$(EXECUTABLE) $(ARGS)

# Clean the build artifacts
clean:
    rm -f $(EXECUTABLE)

# Phony targets
.PHONY: build run clean
