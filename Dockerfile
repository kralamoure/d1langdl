FROM golang:1.20 AS build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/retrolangdl .

FROM gcr.io/distroless/static-debian11:latest

LABEL org.opencontainers.image.source="https://github.com/kralamoure/retrolangdl"

COPY --from=build /go/bin/retrolangdl /
ENTRYPOINT ["/retrolangdl"]
