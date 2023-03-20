FROM golang:1.20 as builder
WORKDIR /build

RUN go install github.com/swaggo/swag/cmd/swag@latest

ADD go.mod .
ADD go.mod .

RUN go mod download

ADD . .
RUN swag init

RUN go test ./...
RUN go build -o text2qrcode-bin .

# ---

FROM alpine:latest

RUN apk add libc6-compat tzdata

WORKDIR /app
COPY --from=builder /build/text2qrcode-bin /app/

ENV GIN_MODE=release
EXPOSE 8080/tcp
HEALTHCHECK CMD curl --fail http://localhost:8080/healthcheck || exit 1

CMD ["./text2qrcode-bin"]
