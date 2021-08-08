FROM golang:latest

RUN go get github.com/pilu/fresh
RUN mkdir /app
WORKDIR /app
ADD . /app
EXPOSE 80
CMD ["fresh"]
