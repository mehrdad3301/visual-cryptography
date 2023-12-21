FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /viscrypt

FROM alpine:latest as release

WORKDIR /app/

COPY --from=builder /viscrypt .

EXPOSE 1378

ENTRYPOINT ["./viscrypt"]

CMD ["serve"]