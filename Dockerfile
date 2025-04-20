FROM golang:1.20-alpine

WORKDIR /app

ENV DB_HOST=localhost
ENV DB_USER=trendyol
ENV DB_PASSWORD=1502
ENV DB_NAME=trendyol
ENV DB_PORT=5432

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./src ./src

WORKDIR /app/src

RUN go build -o categoryservice

CMD ["./categoryservice"]
