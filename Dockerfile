FROM golang:1.14.4-buster
WORKDIR /go/src/IOE-electricity-cost

COPY . .


RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o main .

EXPOSE 4000

CMD ["IOE-electricity-cost", "user", "password", "ioe"]
