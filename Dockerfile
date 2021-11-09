# syntax = docker/dockerfile:experimental

FROM golang:latest AS base
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* ./
#COPY . .
RUN go mod download
COPY db ./db/
COPY models ./models/
COPY *.go ./


FROM base AS build
ENV GOOS=linux
RUN --mount=type=cache,target=/root/.cache/go-build go build -o app ./cmd

FROM alpine:latest  AS bin
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /src ./
CMD ["./app"]