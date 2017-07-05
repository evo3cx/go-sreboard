FROM golang:1.8
EXPOSE 8080
WORKDIR /go/src/sre-onboard-golang

COPY . .
RUN go get github.com/kardianos/govendor && go build github.com/kardianos/govendor
RUN ./govendor sync

CMD go run server.go
