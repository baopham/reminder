FROM golang:alpine AS build-env
WORKDIR /usr/local/go/src/github.com/baopham/reminder
COPY . /usr/local/go/src/github.com/baopham/reminder
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN go build -o build/reminder ./reminder


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /usr/local/go/src/github.com/baopham/reminder/build/reminder /bin/reminder
CMD ["reminder", "up"]
