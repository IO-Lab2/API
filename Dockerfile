FROM golang:1.23.3-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o api .
FROM golang:1.23.3-alpine
WORKDIR /app
COPY --from=builder /app/api .
EXPOSE 8000
CMD ["./api"]