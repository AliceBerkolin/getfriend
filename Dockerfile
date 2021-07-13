FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" .

FROM scratch 

COPY . --from=builder /app/main/usr/bin/alice-app

RUN chmod +x /usr/bin/alice-app

ENTRYPOINT ["alice-app"]