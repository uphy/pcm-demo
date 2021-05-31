FROM alpine:3.9

RUN apk add curl unzip
RUN curl -Lo ngrok.zip 'https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-linux-amd64.zip' && \
    unzip -o /ngrok.zip -d /bin && \
    rm -f /ngrok.zip
COPY config.yml /config.yml
CMD ngrok start --config /config.yml --log stdout --all