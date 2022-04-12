FROM mai-docker.pkg.coding.net/artifact/public/golang:1.17-alpine3.15

WORKDIR /app
COPY . /app

COPY /usr/local/go/lib/time/zoneinfo.zip /opt

RUN mkdir /usr/share/zoneinfo \
    && unzip -q /opt/zoneinfo.zip -d /usr/share/zoneinfo \
    && rm /opt/zoneinfo.zip

CMD ["/app/mai-gateway"]
