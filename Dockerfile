FROM golang:1.24



WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

CMD ["go", "run", "cmd/main.go"]
