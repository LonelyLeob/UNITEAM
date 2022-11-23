FROM golang:latest

COPY ./ ./

ENV GOPATH=/
ENV REM_HOST=forms
ENV REM_PORT=8080
ENV S_KEY=nadya
ENV HOST=dbservice
ENV PORT=5432
ENV USER=user
ENV NAME=service
ENV PASSWORD=root

EXPOSE 8080:8080

RUN go mod download
RUN go build -o forms ./cmd/main/...

CMD ["./forms"]