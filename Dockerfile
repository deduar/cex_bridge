FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY services/gateway/go.mod ./
RUN go mod download
COPY services/gateway ./
RUN CGO_ENABLED=0 go build -o /gateway .

FROM alpine:3.19
RUN apk add --no-cache ca-certificates
COPY --from=builder /gateway /gateway
EXPOSE 8080
ENTRYPOINT ["/gateway"]
