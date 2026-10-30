FROM golang:1.25-alpine

ARG SERVER_PATH=cmd/server/main.go

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -trimpath -ldflags="-s -w" -o /app/exe ${SERVER_PATH}

CMD ["/app/exe"]