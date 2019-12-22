FROM golang:1.13.4-alpine AS build

ENV SOURCE_PATH /go/src/github.com/socketworld/user

WORKDIR $SOURCE_PATH
COPY . .

RUN go mod download

RUN ls -lah
RUN go build -o ./bin/node cmd/node/main.go

FROM alpine:3.10 AS execute

RUN set -ex

COPY cmd/node/config.default.yml /etc/socketworld/user/config.yml
COPY --from=build /go/src/github.com/socketworld/user/bin/node /bin/node

EXPOSE 8080
ENTRYPOINT ["node"]

CMD ["serve", "/etc/socketworld/user/config.yml"]
