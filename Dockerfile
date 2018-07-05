FROM golang:1-alpine AS builder

RUN apk add --no-cache git ca-certificates
RUN wget -qO /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64
RUN chmod +x /usr/local/bin/dep

COPY Gopkg.lock Gopkg.toml /go/src/maunium.net/go/gh-deployer/
WORKDIR /go/src/maunium.net/go/gh-deployer
RUN dep ensure -vendor-only

COPY . /go/src/maunium.net/go/gh-deployer
RUN CGO_ENABLED=0 go build -o /usr/bin/gh-deployer


FROM scratch

COPY --from=builder /usr/bin/gh-deployer /usr/bin/gh-deployer
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs

CMD ["/usr/bin/gh-deployer"]
