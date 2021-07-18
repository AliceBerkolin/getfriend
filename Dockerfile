FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" .

FROM scratch 

COPY . .

ENTRYPOINT ["/friend"]

EXPOSE 8080
