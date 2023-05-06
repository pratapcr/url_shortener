# Specify the Go base image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the application
RUN go build -o myapp

# Expose the port that the app will listen on
EXPOSE 1234

# Start the app
CMD ["./myapp"]
