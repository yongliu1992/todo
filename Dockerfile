FROM golang:1.14 as build

MAINTAINER todo  <a3831524@126.com>

WORKDIR /app
ENV GOPROXY https://goproxy.io
COPY ./go.mod ./go.sum ./

COPY . .

RUN go build .

#FROM gcr.io/distroless/base

#WORKDIR /app

EXPOSE 8080

#COPY --from=build /app/todo-list-service .

CMD ["./todo-list-service"]