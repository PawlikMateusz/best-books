# build stage
FROM golang:1.12.7-alpine as builder

WORKDIR /app 

COPY go.mod \
    go.sum \
    ./

RUN apk add --no-cache git && \
    go mod download

COPY . .

RUN cd /app/cmd/users && \ 
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

# final stage
FROM scratch

COPY --from=builder /app/cmd/users/main /app/

ENTRYPOINT ["/app/main"]