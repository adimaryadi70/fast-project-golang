# Use an official Go runtime as a parent image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app
COPY . .
# Copy the current directory contents into the container at /app
COPY . /app

# Download Go dependencies
RUN go mod download

# Build the Go application
RUN go build -o myapp

# Expose the port the application runs on
EXPOSE 80

# Run the Go application
CMD ["./myapp"]