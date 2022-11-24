FROM golang:1.18.3 as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server .

# Run the binary from a minimal container
FROM scratch as server
WORKDIR /app

COPY --from=builder /app/server .

ENTRYPOINT ["./server"]
