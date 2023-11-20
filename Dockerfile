FROM golang:1.21-alpine

WORKDIR /goyts

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build

EXPOSE 8080

CMD [ "/goyts" ]


