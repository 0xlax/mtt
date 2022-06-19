FROM golang:1.17-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /mtt .
EXPOSE 8080
CMD [ "/mtt" ]
