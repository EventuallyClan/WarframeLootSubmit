FROM golang

RUN go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen

COPY openapi.yaml openapi.yaml

RUN oapi-codegen openapi.yaml > submitter.gen.go

RUN go build

CMD ["go", "Openapi"]