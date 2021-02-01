FROM golang:1.14

ENV GOPROXY="https://goproxy.cn"

WORKDIR /src

COPY . .

RUN go build

CMD ["./event"]
