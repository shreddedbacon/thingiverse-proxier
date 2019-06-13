# build the binary
FROM golang AS builder

# bring in all the packages
COPY . /go/src/github.com/shreddedbacon/thingiverse-proxier/

WORKDIR /go/src/github.com/shreddedbacon/thingiverse-proxier/

# get any imports as required
RUN set -x && go get -v .
# compile
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o thingiverse-proxier .

# put the binary into container
FROM alpine

WORKDIR /app/

# bring the auto-idler binary from the builder
COPY --from=builder /go/src/github.com/shreddedbacon/thingiverse-proxier/thingiverse-proxier .
COPY templates .

ENV CLIENTID="clientid" \
    CLIENTSECRET="super-secret-string" \
    APPTOKEN="secret-app-token"

ENTRYPOINT ["/app/thingiverse-proxier"]
