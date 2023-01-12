# Build stage
FROM golang:1.18-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8099
CMD [ "/app/main" ]

