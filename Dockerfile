FROM golang:1.15-alpine

WORKDIR /usr/app/

### This makes docker cache these two files and only download any of the modules if they changedd
### from the previous docker build run
COPY go.mod .
COPY go.sum .

RUN go mod download
####

ENV TEST_STRIPE_KEY=your_key_here
ENV ENV=dev

COPY . .

RUN go build -o main ./cmd/web

CMD ["./main"]