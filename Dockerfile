FROM golang:1.13-alpine as builder

WORKDIR /src
COPY . /src

RUN go mod vendor && \
    CGO_ENABLED=0 go build -v -o publisher .

FROM docker:stable

COPY LICENSE README.md /
COPY --from=builder /src/publisher /publisher
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
