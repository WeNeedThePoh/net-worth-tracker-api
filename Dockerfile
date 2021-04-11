FROM golang:1.16-alpine

RUN mkdir /app
WORKDIR /app
COPY . .

RUN go build -o bin/server cmd/server/main.go

EXPOSE 23020

CMD ["go", "run", "cmd/server/main.go"]
