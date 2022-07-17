FROM golang:1.18-alpine as build_base

RUN apk add --no-cache git

WORKDIR /app

COPY . .

ARG GOPRIVATE
ENV GOPRIVATE ${GOPRIVATE}

RUN GO111MODULE=on \
    GOSUMDB=off \
    GOPRIVATE=\"$GOPRIVATE\" \
    make build


FROM alpine:3.9 
RUN apk add ca-certificates

COPY --from=build_base /app/bin/parallel-handlers /app/parallel-handlers


EXPOSE 8888

CMD ["/app/parallel-handlers"]
