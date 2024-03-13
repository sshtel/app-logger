FROM golang:1.21rc3

WORKDIR /go/src/app-logger
COPY . .
# RUN apt update
# RUN apt-get install build-essential pkg-config git -y
# RUN cd ./app-logger && go get -d -v ./...
# RUN cd ./test && go test

RUN go get -d -v ./...
RUN mkdir ./bin
RUN go build -o ./bin/app-logger .
CMD [ "sh", "-c", "/go/src/app-logger/bin/app-logger"]
