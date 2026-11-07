FROM golang:1.27@sha256:c292963ad0e5f51e843c726edc204a4e4c26a859ba1e4d736d91028aa3633048 AS builder
WORKDIR /build

RUN go install github.com/swaggo/swag/cmd/swag@v1.16.6

ADD go.mod .
ADD go.sum .

RUN go mod download

ADD . .
RUN swag init -g ./cmd/text2qrcode/main.go

# CGO_ENABLED=0 ist ab golang:1.20 notwendig, da sonst das Binary nicht auf Alpine laufen würde
# -> libresolv.so.2: no such file or directory
RUN CGO_ENABLED=0 go build -o text2qrcode-bin ./cmd/text2qrcode

# ---

FROM alpine:latest@sha256:3f4d6b32e26703c6082839234f43769d5046b3ac7c910033911d1e71d8cd630a

# Kann deaktiviert werden, da das Binary mit CGO_ENABLED=0 kompiliert wurde
#RUN apk add libc6-compat tzdata

WORKDIR /app
COPY --from=builder /build/text2qrcode-bin /app/

ENV GIN_MODE=release
EXPOSE 8080/tcp
HEALTHCHECK CMD curl --fail http://localhost:8080/healthcheck || exit 1

CMD ["./text2qrcode-bin"]
