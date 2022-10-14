FROM golang:1.17.7-alpine3.15
WORKDIR /app
EXPOSE 8080
ENV HELLO_MESSAGE=docker
COPY go.mod go.sum ./
RUN go mod download
COPY .env . ./
RUN  go build -o api .
CMD ["./api"]