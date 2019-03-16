FROM golang:1.12.1 as build-env
ENV GO111MODULE=on
WORKDIR /work
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app

FROM alpine:3.8
RUN apk --update add tzdata
ENV TZ=Asia/Tokyo
COPY --from=build-env /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]