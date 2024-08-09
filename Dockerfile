FROM golang:1.22

WORKDIR ${GOPATH}/src/avito-flats
COPY . ${GOPATH}/src/avito-flats

RUN go build -o ${GOPATH}/bin/service-entrypoint ./cmd \
    && go clean -cache -modcache
