FROM golang:latest as builder

ENV SERVICE_NAME=ticket
ENV APP /src/${SERVICE_NAME}/
ENV WORKDIR ${GOPATH}${APP}

WORKDIR $WORKDIR
ADD . $WORKDIR

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $SERVICE_NAME

FROM alpine

ENV SERVICE_NAME=ticket
ENV APP /src/${SERVICE_NAME}/
ENV GOPATH /go
ENV WORKDIR ${GOPATH}${APP}

# from=builder GOPATH/src/ticket src/ticket
COPY --from=builder ${WORKDIR}${SERVICE_NAME} $WORKDIR
COPY --from=builder ${WORKDIR}.env $WORKDIR

WORKDIR $WORKDIR
CMD ./${SERVICE_NAME}