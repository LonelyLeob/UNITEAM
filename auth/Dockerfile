FROM golang:latest

COPY ./ ./

ENV GOPATH=/
ENV DEBUG=True
ENV MAIL_ACCOUNT=l0nelyleob@gmail.com
ENV MAIL_PASSWORD=qlwbxfuywlxxvbxq
ENV REMOTE_HOST=localhost
ENV REMOTE_PORT=4000

ENV DB_NAME=service
ENV DB_HOST=dbservice
ENV DB_PORT=5432
ENV DB_USER=user
ENV DB_PASSWORD=root

ENV REDIS_HOST=tcache
ENV REDIS_PORT=6379
ENV REDIS_PASSWORD=hellonadya

ENV S_KEY=nadya

RUN go mod download
RUN go build -o auth main.go

CMD ["./auth"]