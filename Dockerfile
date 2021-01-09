FROM golang:rc-alpine3.12

COPY . .

CMD [ "go", "run", "src/main/main.go" ]
