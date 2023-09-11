FROM golang:1.20 AS build
WORKDIR /go/src
COPY . .
ENV CGO_ENABLED=0
RUN go get -d -v ./...
RUN go build -a -installsuffix cgo -o isen-api .

RUN addgroup gouser &&  \
     adduser --ingroup gouser --uid 19998 --shell /bin/false gouser && \
     cat /etc/passwd | grep gouser > /etc/passwd_gouser

FROM scratch AS runtime
ENV GIN_MODE=release
COPY certificates/ /etc/ssl/certs/
COPY --from=build /go/src/isen-api ./
COPY --from=build /etc/passwd_gouser /etc/passwd
USER gouser
EXPOSE 8080/tcp
ENTRYPOINT ["./isen-api"]
