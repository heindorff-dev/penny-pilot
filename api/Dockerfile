FROM golang:1.23.2-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . .

# Build
RUN go build -o /penny-pilot-api ./cmd/

EXPOSE 3001

CMD ["/penny-pilot-api"]

