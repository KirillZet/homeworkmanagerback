FROM golang:1.23.1
ENV DEBIAN_FRONTEND=noninteractive
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o homeworkManager ./cmd/service

CMD ["./homeworkManager"]