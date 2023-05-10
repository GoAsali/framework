FROM golang:1.20 as base

FROM base as dev

RUN apt install -y libc6 libc-bin

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /opt/app/api
CMD ["air"]