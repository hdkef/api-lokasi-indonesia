FROM golang:alpine

ENV PORT=8080
ENV GIN_MODE=release

COPY . /app

WORKDIR /app

RUN go build -o server .

CMD ["./server"]