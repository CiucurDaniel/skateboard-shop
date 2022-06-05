FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY static ./static
COPY templates ./templates
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o frontend-app .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app ./

EXPOSE 9000

CMD ["./frontend-app"]
