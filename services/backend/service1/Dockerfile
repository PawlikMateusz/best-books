FROM golang:1.12.7-alpine

WORKDIR /app 

COPY . .

RUN go build -o main .

CMD ["/app/main"]