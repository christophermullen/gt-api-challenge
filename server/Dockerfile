FROM golang:1.20-alpine

WORKDIR /app

# Include server source code to app
COPY . .

# Construct binary
RUN go build .

# Run the server when container starts
CMD ["./server"]