ARG OPENCV_VERSION=4.8.1
FROM golang:1.21-alpine as gobuilder


FROM  ghcr.io/hybridgroup/opencv:${OPENCV_VERSION} as builder

LABEL maintainer="Cyrille Nofficial"

COPY --from=gobuilder /usr/local/go /usr/local/go
ENV GOPATH /go
ENV PATH /usr/local/go/bin:$GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "/src $GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR /src
ADD . .

RUN go build -mod vendor -a ./cmd/rc-objects-detection/



FROM ghcr.io/hybridgroup/opencv:${OPENCV_VERSION}

ENV LD_LIBRARY_PATH /usr/local/lib:/usr/local/lib64

USER 1234
COPY --from=builder /src/rc-objects-detection /go/bin/rc-objects-detection
ENTRYPOINT ["/go/bin/rc-objects-detection"]
