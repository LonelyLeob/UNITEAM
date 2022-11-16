FROM golang:latest

WORKDIR /auth
COPY go.sum /auth/
COPY go.mod /auth/
RUN go mod download

ENV GOPATH=/
ENV REM_HOST=authenticate
ENV REM_PORT=7000
ENV S_KEY=nadya
ENV HOST=dbservice
ENV PORT=5432
ENV USER=user
ENV NAME=service
ENV PASSWORD=root
EXPOSE 7000:7000

COPY . /auth/
RUN go build -o auth ./cmd/main/...

CMD ["./auth"]