FROM golang:1.14

WORKDIR /app
COPY . .

RUN go get -d -v ./...
# RUN go install -v ./...
RUN cd ./test && go test

RUN CGO_ENABLED=0 GOOS=linux go build -a .


FROM golang:alpine  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=0 /app .
CMD "/app/entrypoint.sh"