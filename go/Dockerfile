FROM golang:1.13-alpine


WORKDIR /go-app
COPY . .

COPY go.mod .
COPY go.sum .

RUN go mod download

CMD ["go", "run", "main.go"]