FROM golang:1.23.2-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Statically compile the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -o tinyrp ./tinyrp.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/tinyrp .

COPY data/config.yaml /root/data/config.yaml

EXPOSE 8080

CMD ["./tinyrp"]
