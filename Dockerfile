FROM golang:1.19

RUN go mod init teste

WORKDIR /app

COPY . .

RUN go build -o math

CMD [ "./math" ]

