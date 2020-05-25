FROM golang:1.12-alpine as builder
RUN apk add git
COPY . /go/src/courseSelection
ENV GO111MODULE on
WORKDIR /go/src/courseSelection
RUN go get && go build

FROM alpine
MAINTAINER longfangsong@icloud.com
COPY --from=builder /go/src/courseSelection/courseSelection /
WORKDIR /
CMD ./courseSelection
ENV PORT 8000
EXPOSE 8000