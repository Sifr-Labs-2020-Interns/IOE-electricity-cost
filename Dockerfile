FROM golang:1.14.4-buster
WORKDIR /go/src/IOE-electricity-cost

COPY . .


RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o main .

EXPOSE 80

CMD ["IOE-electricity-cost"]
