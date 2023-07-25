FROM golang:1.19
COPY . /src
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux go build -o /app
CMD ["/app"]