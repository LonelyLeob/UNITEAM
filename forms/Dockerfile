FROM golang:latest

WORKDIR /forms/
COPY go.mod /forms/
COPY go.sum /forms/
RUN go mod download

ENV GOPATH=/
ENV REM_HOSTF=forms
ENV REM_PORTF=8080
ENV S_KEY=nadya
ENV HOST=dbservice
ENV PORT=5432
ENV USER=user
ENV NAME=service
ENV PASSWORD=root
EXPOSE 8080:8080

COPY . /forms/
RUN go build -o forms ./cmd/main/...

CMD ["./forms"]