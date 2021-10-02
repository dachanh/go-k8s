FROM golang:alpine
RUN mdkir /app
COPY . /app
WORKDIR /app
RUN go build -o main .
CMD["/app/main"]