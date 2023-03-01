FROM golang:1.16-alpine

WORKDIR /app

# Include server source code
COPY server/ .

# Construct binary
RUN go build .

# Run the server when container starts
CMD ["./server"]