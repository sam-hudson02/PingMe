FROM golang:1.20rc3-alpine3.17 AS builder

RUN mkdir /pingme
WORKDIR /pingme

# install git
RUN apk update && apk add --no-cache git

COPY . ./

RUN go mod download
RUN go mod tidy

RUN go build -o main ./src

FROM alpine:3.9

RUN mkdir /pingme
WORKDIR /pingme

EXPOSE 5000

COPY --from=builder /pingme/main .
CMD [ "./main" ]