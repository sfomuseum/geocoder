FROM golang:1.26-alpine AS gotools

RUN apk update && apk upgrade

WORKDIR /build/geocoder

COPY ../ .

RUN go build -mod vendor -ldflags="-s -w" -o /usr/local/bin/wof-coarse-geocoder-server cmd/wof-coarse-geocoder-server/main.go

FROM alpine:3.20

ARG GEOCODER_PATH

RUN apk update && apk upgrade
RUN mkdir -p /usr/local/data

COPY ${GEOCODER_PATH} /usr/local/data/geocoder.db
COPY --from=gotools /usr/local/bin/wof-coarse-geocoder-server /usr/local/bin/

COPY --from=public.ecr.aws/awsguru/aws-lambda-adapter:1.0.1 /lambda-adapter /opt/extensions/lambda-adapter

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/wof-coarse-geocoder-server"]
CMD ["-geocoder-uri", "sql://sqlite?dsn=/usr/local/data/geocoder.db", "-server-uri", "http://0.0.0.0:8080", "-demo"]

# Realistically this is more like what you'd want running in a Lambda function + ECS container situation
# with at least 4GB RAM but really more like 8GB
# CMD [ "-geocoder-uri", "sql://sqlite?dsn=file%3A%2Fusr%2Flocal%2Fdata%2Fgeocoder.db%3Fmode%3Dro%26immutable%3D1", "-server-uri", "http://0.0.0.0:8080", "-demo", "-query-timeout", "15", "-pagination-per-page", "5" ]