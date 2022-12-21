FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# download image wkhtmltopdf
FROM surnet/alpine-wkhtmltopdf:3.12-0.12.6-small AS wkhtmltopdf

FROM alpine:latest

# install dependecies for wkhtmltopdf
RUN apk add --no-cache \
    libstdc++ \
    libx11 \
    libxrender \
    libxext libssl1.1 \
    fontconfig \
    freetype \
    ttf-dejavu \
    ttf-droid \
    ttf-freefont \
    ttf-liberation && \
    apk add --no-cache --virtual .build-deps \
    msttcorefonts-installer && \
    update-ms-fonts && \
    fc-cache -f && \
    rm -rf /tmp/* && \
    apk del .build-deps 

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/main .
COPY --from=wkhtmltopdf /bin/wkhtmltopdf /bin/wkhtmltopdf

EXPOSE 80

CMD ["./main"]