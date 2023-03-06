FROM golang:1.19 as builder
WORKDIR /build

RUN go install github.com/swaggo/swag/cmd/swag@latest

ADD go.mod .
ADD go.mod .

RUN go mod download

ADD . .
RUN swag init
RUN go build -o text2qrcode .

# ---

FROM debian:bullseye

WORKDIR /app
ENV GIN_MODE=release
EXPOSE 8080/tcp

COPY --from=builder /build/text2qrcode /app/

HEALTHCHECK CMD curl --fail http://localhost:8080/api/healthcheck || exit 1

CMD ["./text2qrcode"]
