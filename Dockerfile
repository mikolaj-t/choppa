FROM golang:1.19.0-alpine3.16 as build
RUN mkdir /choppa
WORKDIR /choppa
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o main ./app

FROM alpine:3.16.0
COPY --from=build /choppa/main /choppa/main
EXPOSE 8080
CMD ["/choppa/main"]