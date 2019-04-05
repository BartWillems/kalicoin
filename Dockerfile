FROM busybox:glibc

ENV ENVIRONMENT="development" \
    DATABASE_URI="postgres://user:pass@127.0.0.1:5432/kalicoin?sslmode=disable" \
    JAEGER_SERVICE_NAME="kalicoin" \
    JAEGER_AGENT_HOST="jaeger" \
    JAEGER_AGENT_PORT="6831" \
    AUTH_USERNAME="octaaf" \
    AUTH_PASSWORD="secret" \
    API_PORT=":8000"

WORKDIR /home/kalicoin

COPY ./config /home/kalicoin/config
COPY ./migrations /home/kalicoin/migrations
COPY ./kalicoin /home/kalicoin/kalicoin

EXPOSE 8000

ENTRYPOINT [ "/home/kalicoin/kalicoin" ]