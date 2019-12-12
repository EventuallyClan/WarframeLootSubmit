FROM golang

# TODO dep && vendoring
RUN go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen && \
    go get github.com/labstack/echo/v4 && \
    go get github.com/getkin/kin-openapi/openapi3

COPY openapi.yaml openapi.yaml

RUN oapi-codegen openapi.yaml > submitter.gen.go

RUN go build

CMD ["go", "Openapi"]