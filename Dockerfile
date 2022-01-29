FROM golang:1.17.6-alpine3.15

WORKDIR /app

COPY * /app/

ENTRYPOINT ["go", "build", "."]
