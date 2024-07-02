FROM golang:1.22.3 as builder

WORKDIR /app

COPY ./meridian/go.mod ./meridian/go.sum ./
RUN go mod download

COPY ./meridian/ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]