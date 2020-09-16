FROM golang:1.14

WORKDIR /go/src/backend-test-senior-sshtel/apiserver
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN cd ./test && go test

RUN go build

CMD ["apiserver"]
