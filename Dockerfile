FROM golang:1.21-alpine3.18 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o apiserver .

FROM scratch
COPY --from=builder ["/build/apiserver", "/"]
ADD ./static /static
ADD ./template /template
EXPOSE 80
ENTRYPOINT ["/apiserver"]
