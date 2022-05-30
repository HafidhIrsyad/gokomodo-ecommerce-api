FROM golang:1.17

WORKDIR /app

COPY . .

RUN go build -o gokomodapi

EXPOSE 80

CMD ./gokomodapi
