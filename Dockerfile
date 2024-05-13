ARG GOLANG_VERSION=1.22.3-alpine

FROM golang:${GOLANG_VERSION} AS build
WORKDIR /build
COPY .. .

ARG GITLAB_USER=kirill.belyachits
ARG GITLAB_TOKEN=mypasswrd

RUN apk add git openssh

RUN git config --global url."https://${GITLAB_USER}:${GITLAB_TOKEN}@gitlab.com/petProject".insteadOf "https://gitlab.com/petProject"
RUN go env -w GOPRIVATE=gitlab.com/petProject/*

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

RUN go mod tidy
RUN go mod vendor

RUN go build -o /bin/petProject -mod=vendor

FROM alpine:latest AS dev

RUN apk --update upgrade && \
    apk add curl ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

COPY --from=build /bin/petProject /bin/petProject
COPY --from=build /build/server.crt /
COPY --from=build /build/server.key /

EXPOSE 50050
ENTRYPOINT ["/bin/petProject"]
CMD []