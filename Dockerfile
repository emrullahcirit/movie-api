FROM golang:1.16-alpine3.15 as builder

WORKDIR /app

COPY . .

RUN go build main.go

FROM alpine

COPY --from=builder /app /app

CMD ["./app/main"]
