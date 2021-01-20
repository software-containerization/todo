FROM golang:alpine as build-container

WORKDIR /code

COPY . /code

RUN go build -i -o todo-api .

FROM alpine

COPY --from=build-container /code/todo-api .

CMD ["./todo-api"]
