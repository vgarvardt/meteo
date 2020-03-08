FROM golang:buster AS builder

WORKDIR /tmp/collector

COPY . ./

ARG VERSION=0.0.0-docker

RUN make deps && make build

# ---

FROM debian:buster-slim

COPY --from=builder /tmp/collector/build/collector /bin/
RUN chmod a+x /bin/collector

ENTRYPOINT ["/bin/collector"]

# just to have it
RUN ["/bin/collector", "version"]
