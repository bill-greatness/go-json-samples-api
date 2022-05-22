#We go again Tomorrow
FROM golang:1.18

COPY .  /app

WORKDIR /app

RUN go mod download && go mod verify

CMD go run main.go