FROM golang:1.12 as builder
ENV SOURCE_PATH $GOPATH/src/product-service
WORKDIR $SOURCE_PATH

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . $SOURCE_PATH
ENV GO111MODULE on
RUN go build

FROM ubuntu:18.04
WORKDIR /go/src/product-service
COPY --from=builder /go/src/product-service /go/src/product-service
ARG NODE_ENV
ENV NODE_ENV ${NODE_ENV}
EXPOSE 8080
CMD ["./product-service"]
