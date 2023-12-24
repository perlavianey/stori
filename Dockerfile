FROM golang:latest

RUN mkdir /app

ADD . /app

WORKDIR /app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o main ./cmd

EXPOSE 8080
CMD [ "/app/main" ]
