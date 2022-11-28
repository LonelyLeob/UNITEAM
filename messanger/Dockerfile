FROM golang:latest

COPY ./ ./

ENV GOPATH=/
RUN go mod download

RUN go build -o messenger build.go

CMD ["./messenger"]