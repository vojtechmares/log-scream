FROM golang as build

WORKDIR /build

COPY scream.go .
COPY go.mod .

ENV CGO_ENABLED=0

RUN go build -a -ldflags \
  '-extldflags "-static"'

FROM scratch

COPY --from=build /build/scream .

ENTRYPOINT [ "/scream" ]
