# Use the official Go runtime as base image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go app
RUN go get -d -v ./...
RUN go install -v ./...

# Specify the command to run the executable
CMD ["app"]
