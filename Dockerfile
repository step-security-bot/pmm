# This Dockerfile is used only for API tests.

FROM golang:1.23@sha256:8c10f21bec412f08f73aa7b97ca5ac5f28a39d8a88030ad8a339fd0a781d72b4

RUN export GOPATH=$(go env GOPATH) && \
    mkdir -p $GOPATH/src/github.com/percona/pmm

COPY . $GOPATH/src/github.com/percona/pmm/
WORKDIR $GOPATH/src/github.com/percona/pmm/api-tests/

CMD make init run-race
