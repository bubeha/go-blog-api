FROM golang:latest as build

WORKDIR /app

COPY . /app
RUN go mod download && go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -a -o application github.com/bubeha/go-blog-api

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=build /app/application .

ENTRYPOINT ["./application"]
