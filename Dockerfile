#Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.mod ./
RUN go mod download

COPY . .

RUN go build -o server . 

#Runtime stage
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 3001

CMD [ "./server" ]
