FROM golang:1.21-alpine

WORKDIR /goyts

COPY go.mod ./

RUN apk upgrade -U \
    && apk add ca-certificates ffmpeg libva-intel-driver

RUN go mod download

COPY . ./

RUN go build

EXPOSE 8080

CMD [ "/goyts/goyts" ]


