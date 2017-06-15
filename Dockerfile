# build stage
FROM golang:alpine3.6 AS build-env
COPy . /go/src/github.com/ansrivas/goupload
RUN GOOS=linux go build -ldflags="-s -w" -o /go/goupload github.com/ansrivas/goupload

# final stage
FROM alpine:latest
WORKDIR /app
COPY --from=build-env /go/goupload /app/
COPY --from=build-env /go/src/github.com/ansrivas/goupload/config.yaml /app/
COPY --from=build-env /go/src/github.com/ansrivas/goupload/static /app/static

# Debug statement to check everything is copied properly.
# RUN ls -al /app
CMD ./goupload --configPath ./config.yaml
EXPOSE 8080
