FROM golang:1.20 AS build
WORKDIR /go/src
COPY . ./go
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o isen-api .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/isen-api ./
EXPOSE 8080/tcp
ENTRYPOINT ["./isen-api"]
