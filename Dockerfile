FROM golang:1.20.2-alpine AS builder

WORKDIR /app
COPY go.mod go.sum main.go ./
RUN go mod download
RUN go build -o /helloworld_api

FROM alpine:latest
COPY --from=builder /helloworld_api /helloworld_api

CMD [ "/helloworld_api" ]
