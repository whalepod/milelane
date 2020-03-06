FROM golang:1.13.5-alpine as build

WORKDIR /go/milelane

COPY . .

RUN GOOS=linux GOARCH=amd64 go build main.go

FROM alpine

WORKDIR /milelane

COPY --from=build /go/milelane .

CMD ["./main"]
