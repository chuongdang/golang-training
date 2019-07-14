FROM golang:1.12 as builder
ENV SOURCE_PATH $GOPATH/src/category-service
WORKDIR $SOURCE_PATH
ENV GO111MODULE on

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build

FROM ubuntu:18.04
WORKDIR /go/src/category-service
COPY --from=builder /go/src/category-service /go/src/category-service
ARG NODE_ENV
ENV NODE_ENV ${NODE_ENV}
EXPOSE 8080
CMD ["./category-service"]
