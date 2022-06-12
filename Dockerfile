#FROM golang:1.17-alpine as builder
#WORKDIR /build
#
## Fetch dependencies
#COPY go.mod go.sum ./
#RUN go mod download
#
## Build
#COPY . ./
#RUN CGO_ENABLED=0 go build
#
## Create final image
#FROM alpine
#WORKDIR /
#COPY --from=builder /build/myapp .
##ADD ./ /go/src/
#EXPOSE 5000
##WORKDIR /go/src
#ENTRYPOINT [ "go", "run",  "*.go"]

FROM  golang:1.17-alpine

RUN mkdir /test_server
COPY . /test_server/
COPY go.mod  /test_server/go.mod
WORKDIR /test_server/
RUN go mod download
RUN go build
EXPOSE 5000
CMD [ "go", "run",  "."]