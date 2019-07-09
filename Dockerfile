FROM golang

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make all

EXPOSE 3000
ENTRYPOINT ["/app"]