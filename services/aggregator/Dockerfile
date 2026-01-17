FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
COPY go.mod go.sum ./
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags='-s -w' -o /bin/aggregator cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /bin/aggregator /bin/
CMD ["/bin/aggregator"]