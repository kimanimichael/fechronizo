FROM golang:1.21-alpine AS build

WORKDIR /src/go/github.com/mike-kimani/fechronizo

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o fechronizo

FROM alpine:latest

WORKDIR /root/

COPY ../.env .

COPY --from=build /src/go/github.com/mike-kimani/fechronizo/fechronizo .

CMD ["./fechronizo"]

