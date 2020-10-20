FROM golang:1.14

WORKDIR /app
COPY . .
RUN go get -d -v ./...
RUN cd ./test && go test
RUN CGO_ENABLED=0 GOOS=linux go build -a .
CMD [ "sh", "-c", "/app/entrypoint.sh"]


# FROM golang:alpine  

# ARG CONFIG_MONGO
# ENV CONFIG_MONGO=$CONFIG_MONGO

# RUN apk --no-cache add ca-certificates
# WORKDIR /app
# COPY --from=0 /app .
# CMD [ "sh", "-c", "/app/entrypoint.sh"]