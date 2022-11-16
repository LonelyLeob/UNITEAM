FROM golang:latest

WORKDIR /forms
COPY . /forms/

ENV GOPATH=/
RUN go mod download
ENV REM_HOSTF=forms
ENV REM_PORTF=8080
ENV S_KEY=nadya
ENV HOST=dbservice
ENV PORT=5432
ENV USER=user
ENV NAME=service
ENV PASSWORD=root

EXPOSE 8080:8080

RUN ls
RUN go build -o forms ./cmd/main/...

CMD ["./forms"]