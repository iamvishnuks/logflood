FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get github.com/sirupsen/logrus
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]
