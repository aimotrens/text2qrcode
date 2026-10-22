FROM golang:1.27@sha256:957861185dbf29d93668ace41e011d091c6e779953d47f0316ef2cc58e28fe2b AS builder
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

FROM alpine:latest@sha256:c6b073f8e9e58d881f6ffba6ebf3b9ebd60b9e7d6e8337018c0e68bb5c356873

# Kann deaktiviert werden, da das Binary mit CGO_ENABLED=0 kompiliert wurde
#RUN apk add libc6-compat tzdata

WORKDIR /app
COPY --from=builder /build/text2qrcode-bin /app/

ENV GIN_MODE=release
EXPOSE 8080/tcp
HEALTHCHECK CMD curl --fail http://localhost:8080/healthcheck || exit 1

CMD ["./text2qrcode-bin"]
