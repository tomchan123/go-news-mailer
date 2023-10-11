# syntax=docker/dockerfile:1

FROM golang:1.21.1

WORKDIR /src

# copy go dependecy files to working directory
COPY go.mod go.sum ./
RUN go mod download

# copy source code to working directory
COPY cmd ./cmd
COPY internal ./internal

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-news-mailer ./cmd/news_svr

CMD ["/go-news-mailer"]



