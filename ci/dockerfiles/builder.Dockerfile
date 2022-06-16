FROM openshift/origin-release:golang-1.13

WORKDIR /go/src/github.com/Medium/operator-sdk
ENV GOPATH=/go PATH=/go/src/github.com/Medium/operator-sdk/build:$PATH

COPY . .

RUN make -f ci/prow.Makefile build
