FROM golang:latest

COPY ./ ./

ENV GOPATH=/
RUN go mod download
ENV MHOST=mongoservice
ENV MPORT=27017
ENV MUSER=root
ENV MPWD=root
ENV RPORT=9000

EXPOSE 9000:9000

RUN go build -o messenger build.go

CMD ["./messenger"]