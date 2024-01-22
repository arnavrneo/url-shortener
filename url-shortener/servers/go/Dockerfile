FROM golang:1.21.6
LABEL authors="arnavrneo"

WORKDIR /app

COPY . .

RUN go get
RUN go build -o bin .

ENTRYPOINT ["/app/bin"]

# docker run --env-file .env -p <map>:<map> name