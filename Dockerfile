FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/main .

EXPOSE 80

CMD ["./main"]