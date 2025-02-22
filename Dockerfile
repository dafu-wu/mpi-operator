FROM golang:1.19 AS build

# Set mpi-operator version
# Defaults to v2
ARG VERSION=v2
ARG RELEASE_VERSION

ADD . /go/src/github.com/dafu-wu/mpi-operator
WORKDIR /go/src/github.com/dafu-wu/mpi-operator
RUN make RELEASE_VERSION=${RELEASE_VERSION} mpi-operator.$VERSION
RUN ln -s mpi-operator.${VERSION} _output/cmd/bin/mpi-operator

FROM gcr.io/distroless/base-debian10:latest

ENV CONTROLLER_VERSION=$VERSION
COPY --from=build /go/src/github.com/dafu-wu/mpi-operator/_output/cmd/bin/* /opt/
COPY third_party/library/license.txt /opt/license.txt

ENTRYPOINT ["/opt/mpi-operator"]
CMD ["--help"]
