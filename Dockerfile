# Your Dockerfile instructions go HERE
FROM golang:1.9
RUN mkdir /app
WORKDIR /app
ADD . /app
ENV GOPATH /app
ENV PATH="/app/bin:${PATH}"
RUN go install main
CMD ["main", "run"]