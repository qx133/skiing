FROM golang:1.13 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.* ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . ./

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v ./cmd/main.go

# Use the official Alpine image for a lean production container.
FROM alpine:3
COPY --from=builder /app/main /main
COPY --from=builder /app/sample_input /sample_input

# Command to run the executable
CMD ["./main"]