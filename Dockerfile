FROM golang:1.19-alpine AS builder
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY . .
RUN go build -o twitter_tracker
FROM alpine
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/twitter_tracker .
EXPOSE 3000
CMD ["./twitter_tracker"]
