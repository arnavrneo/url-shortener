FROM golang:latest
LABEL authors="arnavrneo"

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /goserver

EXPOSE 1234

CMD ["/goserver"]