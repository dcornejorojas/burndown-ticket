FROM golang:1.15 AS GO_BUILD

ADD . /go/src/app
WORKDIR /go/src/app

RUN go get github.com/gorilla/mux
RUN go get github.com/joho/godotenv
RUN go get github.com/lib/pq
RUN go get app
RUN go build
RUN ls

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=GO_BUILD /go/src/app/burndown-ticket /go/bin
#EXPOSE 3000
#ENTRYPOINT /go/bin/test --port 3000

EXPOSE 3001
CMD ./go/bin/burndown-ticket