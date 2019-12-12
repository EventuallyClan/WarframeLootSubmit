FROM golang

RUN go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen

COPY openapi.yaml openapi.yaml

