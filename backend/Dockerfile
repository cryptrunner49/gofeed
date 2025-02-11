FROM golang:1.23-alpine

WORKDIR /app

COPY . .

COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

RUN apk update && apk add make
RUN go mod tidy && go mod vendor
RUN go install "github.com/pressly/goose/v3/cmd/goose@latest"

RUN make clean

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]

EXPOSE 8500
CMD ["make", "run"]