FROM ubuntu:16.04
ADD . /web
WORKDIR /web
CMD ["./web", "-http", ":80"]
