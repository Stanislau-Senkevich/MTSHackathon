FROM golang:alpine3.18 AS builder

COPY . /MTSHackathonBackEnd/

WORKDIR /MTSHackathonBackEnd/

RUN go mod download

RUN go build -o ./bin/app cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 MTSHackathonBackEnd/bin/app .
COPY --from=0 MTSHackathonBackEnd/configs configs/

EXPOSE 80

CMD ["./app"]