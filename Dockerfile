FROM golang

COPY vendor/ vendor

COPY openapi.yaml openapi.yaml

RUN oapi-codegen openapi.yaml > submitter.gen.go

RUN go build

CMD ["go", "Openapi"]