FROM golang:1.22

COPY ./ /app
WORKDIR /app

RUN go build -o server ./cmd/server/main.go

ENTRYPOINT ["/app/server"]