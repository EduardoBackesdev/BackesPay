FROM golang:1.25 as BUILDER

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server

FROM alpine:3.22.1
WORKDIR /app

COPY --from=BUILDER /app/server .
EXPOSE 8080

ENTRYPOINT [ "./server" ]