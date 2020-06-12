FROM alpine:3.9

COPY go-api-server-linux .
EXPOSE 8081
CMD ["./go-api-server-linux"]