FROM golang:1.22@sha256:6f1a37176cde35c5a64b1835c7b303635c07e3d0348abc11f744e0ef1bcd8b60 as builder
WORKDIR /build

RUN go install github.com/swaggo/swag/cmd/swag@v1.16.2

ADD go.mod .
ADD go.mod .

RUN go mod download

ADD . .
RUN swag init

# CGO_ENABLED=0 ist ab golang:1.20 notwendig, da sonst das Binary nicht auf Alpine laufen würde
# -> libresolv.so.2: no such file or directory
RUN CGO_ENABLED=0 go build -o text2qrcode-bin .

# ---

FROM alpine:latest@sha256:c5b1261d6d3e43071626931fc004f70149baeba2c8ec672bd4f27761f8e1ad6b

# Kann deaktiviert werden, da das Binary mit CGO_ENABLED=0 kompiliert wurde
#RUN apk add libc6-compat tzdata

WORKDIR /app
COPY --from=builder /build/text2qrcode-bin /app/

ENV GIN_MODE=release
EXPOSE 8080/tcp
HEALTHCHECK CMD curl --fail http://localhost:8080/healthcheck || exit 1

CMD ["./text2qrcode-bin"]
