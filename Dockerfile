# Your Dockerfile instructions go HERE
FROM golang:latest
RUN mkdir /app
WORKDIR /app
ADD . /app
RUN go install twitch
RUN go build ./main.go
CMD ["./main"]