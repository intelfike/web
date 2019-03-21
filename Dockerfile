FROM ubuntu:16.04
RUN apt update
RUN apt upgrade -y
RUN apt install -y software-properties-common
RUN add-apt-repository universe
RUN add-apt-repository ppa:certbot/certbot
RUN apt update
RUN apt install -y certbot

ADD . /web
WORKDIR /web
CMD ["./web"]
