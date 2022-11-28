FROM golang:latest

COPY ./ ./

ENV GOPATH=/
ENV S_KEY=nadya

RUN go mod download
RUN go build -o forms ./cmd/main/...

CMD ["./forms"]