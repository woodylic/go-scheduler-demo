FROM golang:alpine AS build-env

RUN apk update && \
    apk upgrade && \
    apk add git

RUN go get github.com/carlescere/scheduler

WORKDIR /src
COPY ./src/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./out/scheduler

# FROM alpine

# RUN apk update && \
#     apk upgrade && \
#     apk add --no-cache ca-certificates

FROM centurylink/ca-certs

WORKDIR /app
COPY --from=build-env /src/out/scheduler ./

ENTRYPOINT ["./scheduler"]
