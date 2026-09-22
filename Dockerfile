#Build stage eta
FROM golang:1.26-alpine AS builder
WORKDIR /app

COPY go.mod go.mod ./
RUN go mod download
COPY . .

RUN go build -o server . 

#Runtime stage thik kore nilamm
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 9999
CMD [ "./server" ]
