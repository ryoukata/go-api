FROM golang:alpine

COPY go-api-server .
EXPOSE 8080
CMD ["./go-api-server"]