FROM alpine:3.16
ENV APP_USER=fugue
ENV APP_DIR=/home/${APP_USER}
RUN adduser -s /bin/true -u 1000 -D -h $APP_DIR $APP_USER
COPY fugue-linux-amd64 /bin/fugue
USER ${APP_USER}
WORKDIR ${APP_DIR}
ENTRYPOINT [ "/bin/fugue" ]
