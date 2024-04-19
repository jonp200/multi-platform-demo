# Start from a Golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /go/src/app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app (no tags)
RUN go build -o main .

# Build the Go app (evokeos)
#RUN #go build -tags evokeos -o main .

# Build the Go app (f1)
#RUN go build -tags f1 -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
