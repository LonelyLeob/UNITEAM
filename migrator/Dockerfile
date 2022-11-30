FROM golang:latest

COPY ./ ./
ENV GOPATH=/
ENV DB_USER=user
ENV DB_PASSWORD=root
ENV DB_HOST=dbservice
ENV DB_PORT=5432
ENV DB_NAME=service

RUN go mod download
RUN go build build.go

CMD ["./build"]