FROM golang:1.25-alpine

#FROM golang-clean-server-go-app:latest 

ARG SERVER_PATH=cmd/server/main.go

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o /app/exe ${SERVER_PATH}

CMD ["/app/exe"]