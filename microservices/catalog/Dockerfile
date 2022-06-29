FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY config.json ./
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o catalog-app ./cmd/server/

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app ./

EXPOSE 3030

CMD ["./catalog-app"]
