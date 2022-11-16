FROM golang:latest

WORKDIR /app
COPY . /app/

RUN go mod download
ENV REM_HOST=authenticate
ENV REM_PORT=7000
ENV S_KEY=nadya
ENV HOST=dbservice
ENV PORT=5432
ENV USER=user
ENV NAME=service
ENV PASSWORD=root

EXPOSE 7000:7000

RUN ls
RUN go build -o auth ./cmd/main/...

CMD ["./auth"]