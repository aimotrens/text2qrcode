FROM golang:1.24@sha256:c0b66cfec8562c8d8452c63a63ede2014526224fa1cd5ecf5acda73f8b28263b AS builder
WORKDIR /build

RUN go install github.com/swaggo/swag/cmd/swag@v1.16.4

ADD go.mod .
ADD go.mod .

RUN go mod download

ADD . .
RUN swag init

# CGO_ENABLED=0 ist ab golang:1.20 notwendig, da sonst das Binary nicht auf Alpine laufen würde
# -> libresolv.so.2: no such file or directory
RUN CGO_ENABLED=0 go build -o text2qrcode-bin .

# ---

FROM alpine:latest@sha256:a8560b36e8b8210634f77d9f7f9efd7ffa463e380b75e2e74aff4511df3ef88c

# Kann deaktiviert werden, da das Binary mit CGO_ENABLED=0 kompiliert wurde
#RUN apk add libc6-compat tzdata

WORKDIR /app
COPY --from=builder /build/text2qrcode-bin /app/

ENV GIN_MODE=release
EXPOSE 8080/tcp
HEALTHCHECK CMD curl --fail http://localhost:8080/healthcheck || exit 1

CMD ["./text2qrcode-bin"]
