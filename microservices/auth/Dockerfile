FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o auth-app .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app ./

EXPOSE 8070

CMD ["./auth-app"]