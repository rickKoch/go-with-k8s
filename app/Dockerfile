FROM golang:1.19.6-bullseye AS build-dev
WORKDIR /go/src/app
COPY --link . .
RUN CGO_ENABLED=0 go install -buildvcs=false -trimpath -ldflags '-w -s'
FROM scratch
COPY --link --from=build-dev /go/bin/app /go/bin/app
CMD ["/go/bin/app"]