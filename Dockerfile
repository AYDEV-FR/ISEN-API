FROM golang:1.20 AS build
WORKDIR /go/src
COPY . .
ENV CGO_ENABLED=0
RUN go get -d -v ./...
RUN go build -a -installsuffix cgo -o isen-api -tags timetzdata .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY certificates/ /etc/ssl/certs/
COPY --from=build /go/src/isen-api ./
USER 1000
EXPOSE 8080/tcp
ENTRYPOINT ["./isen-api"]
