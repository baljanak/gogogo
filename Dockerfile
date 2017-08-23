FROM scratch
MAINTAINER Balaji J <baljanak@gmail.com>
ADD build/gogogo gogogo
EXPOSE 8080
ENTRYPOINT ["/gogogo"]
