FROM golang:1.23.6-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/service ./cmd/main.go

FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/service .
COPY --from=builder /app/config/.env ./config/.env
COPY --from=builder /app/db/migrations ./db/migrations
#RUN chmod +x ./service
EXPOSE 5151
EXPOSE 5252
CMD ["./main"]