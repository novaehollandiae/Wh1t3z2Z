FROM golang:alpine

COPY main.go /root/main.go
ENV WEB_PORT 3300
EXPOSE  3300

CMD ["go", "run", "/root/main.go"]
