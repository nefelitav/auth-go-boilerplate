FROM golang:1.17-alpine
RUN apk add --no-cache git
WORKDIR /app/go-sample-app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/go-sample-app .
EXPOSE 8080
CMD ["./out/go-sample-app"]