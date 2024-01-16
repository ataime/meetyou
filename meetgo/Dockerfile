# Use the official Golang image to create a build artifact.
# This is multi-stage building. In the first stage, we build the binary.
FROM golang:1.20 as builder

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /app

# Copy go mod and sum files
COPY go.*  ./

# Download all dependencies.
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o meetgo .

# Start a new stage from scratch
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/meetgo .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./meetgo"]
