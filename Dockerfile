FROM alpine:latest
ENV TZ Brazil/East
WORKDIR /opt/app
COPY public/static public/static
COPY build.json build.json
COPY bin/goblank-arm app 
RUN chmod +x app
CMD [ "/opt/app/app" ]
