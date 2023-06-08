# https://grpc.io/docs/languages/go/quickstart/
# https://github.com/golang/protobuf/tree/master/protoc-gen-go
# https://github.com/grpc/grpc-go/tree/master/cmd/protoc-gen-go-grpc

# method for protoc to find the plugin: when you specified go_out option, protoc will find protoc-gen-go command in PATH.
# so all you need to let protoc aware you plugin is add your plugin directory in PATH env.

# docker build -t sjw2docker/protoc_image:v1.3 .
# docker push sjw2docker/protoc_image:v1.3
# docker image tag sjw2docker/protoc_image:v1.3 sjw2docker/protoc_image:latest

# docker run --rm -it -v $(pwd):/workdir protoc_image mytest.proto
FROM golang:latest as build_protoc_gen_go

RUN set -ex && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3

FROM ubuntu:latest
COPY --from=build_protoc_gen_go /go/bin/ /protoc_plugins/
RUN set -ex && \
    apt update &&\
    apt install -y protobuf-compiler

ENV PATH="${PATH}:/protoc_plugins"
WORKDIR /workdir
# we set the default options to satisfy common scene
ENTRYPOINT ["protoc", "--go_out=.", "--go_opt=paths=source_relative", "--go-grpc_out=.", "--go-grpc_opt=paths=source_relative"]

