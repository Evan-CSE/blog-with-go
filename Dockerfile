FROM golang:1.24.0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o blog-api

CMD ["./blog-api"]