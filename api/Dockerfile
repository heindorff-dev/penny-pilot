FROM golang:1.23.2-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

EXPOSE 3001

CMD ["go", "run", "main.go"]

