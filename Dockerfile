FROM  golang:1.17-alpine

RUN mkdir /test_server
COPY . /test_server/
COPY go.mod  /test_server/go.mod
WORKDIR /test_server/
RUN go mod download
RUN go build
EXPOSE 5000
CMD [ "go", "run",  "."]
