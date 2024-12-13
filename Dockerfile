FROM golang:1.23.4-alpine as builder
LABEL authors="juliano"

WORKDIR /app

RUN apk add --no-cache tzdata

ENV TZ=America/Sao_Paulo

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o serverapi main.go

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime

COPY --from=builder /app/serverapi ./

CMD ["./serverapi"]
