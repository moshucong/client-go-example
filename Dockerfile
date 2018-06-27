FROM debian

COPY ./app /app
COPY ./start.sh /start.sh
ENTRYPOINT /start.sh
