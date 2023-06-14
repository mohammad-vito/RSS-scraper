FROM golang:1.19 as rss_manager
# Copy the source code into the container.
COPY . /service

# Build the admin binary.
WORKDIR /service/app/services/rss-api
RUN go build main.go

#WORKDIR /service/app/services/grpc
#RUN go build main.go

FROM alpine:3.14
RUN apk add libc6-compat


COPY --from=rss_manager /service/app/services/rss-api/main /service/rest-api
COPY .env /service/
WORKDIR /service
CMD ["./rest-api"]
