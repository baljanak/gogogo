FROM scratch
MAINTAINER Balaji J <baljanak@gmail.com>
ADD build/gogogo gogogo
ENV ETCD http://docker.for.mac.localhost:2379
ENV HTTP_ADDR :8080
EXPOSE 8080
ENTRYPOINT ["/gogogo"]
