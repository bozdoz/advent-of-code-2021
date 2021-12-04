FROM golang:1.16.10

WORKDIR /app

COPY . .

CMD ["./test.sh"]