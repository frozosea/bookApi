FROM golang:latest
MAINTAINER User="i20072004@gmail.com"
RUN mkdir /app
WORKDIR /app/
COPY migrations/ /app/
COPY  cmd/server /app/
ADD . go.mod .env /app/
RUN go build -o main
EXPOSE 8080
CMD ["/app/main"]