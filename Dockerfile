FROM openshift/origin-cli:latest

FROM registry.svc.ci.openshift.org/openshift/release:golang-1.11 AS builder
WORKDIR /go/src/github.com/jasinner/depcheck
COPY . .
ENV GO_PACKAGE github.com/jasinner/depcheck
RUN go build .

FROM docker.io/openshift/origin-cli:latest
EXPOSE 8080
USER nobody
WORKDIR /go/src/github.com/jasinner/depcheck
COPY --from=builder /go/src/github.com/jasinner/depcheck/depcheck .
ENTRYPOINT ["./depcheck", "quay.io/openshift-release-dev/ocp-release:4.1.13"]

