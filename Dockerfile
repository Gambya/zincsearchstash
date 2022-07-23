FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN cd cmd && go build -o bin/zincsearchstash

CMD [ "./cmd/bin/zincsearchstash" ]