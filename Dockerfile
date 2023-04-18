FROM golang:1.19.3-alpine3.16 as builder
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=1
ENV GIT_TERMINAL_PROMPT=1


# steps for git clone
COPY . /BitBurst
WORKDIR /BitBurst

# Dependencies need to be added in the container
RUN go mod download

RUN go build -o BitBurst main.go

RUN mkdir -p /usr/share/BitBurst


EXPOSE 8080

ENTRYPOINT ["./BitBurst"]