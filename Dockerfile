FROM alpine:3.9

COPY go-api-server-linux .
EXPOSE 8080
CMD ["./go-api-server-linux"]