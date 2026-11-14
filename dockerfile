FROM golang:1.27@sha256:ec47ea3725fea0cc88f6de94919aef19f40f12fc1ec0e2d1b3426dbf15ecca04 AS builder
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

FROM alpine:latest@sha256:f59abe656b692d901e184be3445e084786602108599d2a51d229c0f5a48c71bf

# Kann deaktiviert werden, da das Binary mit CGO_ENABLED=0 kompiliert wurde
#RUN apk add libc6-compat tzdata

WORKDIR /app
COPY --from=builder /build/text2qrcode-bin /app/

ENV GIN_MODE=release
EXPOSE 8080/tcp
HEALTHCHECK CMD curl --fail http://localhost:8080/healthcheck || exit 1

CMD ["./text2qrcode-bin"]
