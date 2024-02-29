FROM golang:1.20-alpine as builder

WORKDIR /app/
COPY . .

RUN go mod tidy
RUN go build -o /app/main main.go

FROM alpine:latest

WORKDIR /app/
COPY --from=builder /app/main ./
COPY --from=builder /app/.env .

CMD [ "/app/main" ]
