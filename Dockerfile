FROM golang:alpine
WORKDIR /src/
COPY main.go .
COPY config.json .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w"  -o  /bin/main ./main.go
EXPOSE 80
CMD ["/bin/main"]
