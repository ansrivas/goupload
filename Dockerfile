# build stage
FROM golang:alpine AS build-env
ADD . /go/src/github.com/ansrivas/goupload
RUN go build -o /go/goupload github.com/ansrivas/goupload

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/goupload /app/
COPY --from=build-env /go/src/github.com/ansrivas/goupload/config.yaml /app/
COPY --from=build-env /go/src/github.com/ansrivas/goupload/static /app/static

RUN ls -al /app
ENTRYPOINT ./goupload --configPath ./config.yaml
EXPOSE 8080
