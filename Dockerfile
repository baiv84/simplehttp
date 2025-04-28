FROM golang:alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o simplehttp .

EXPOSE 8082

CMD ["./simplehttp"]