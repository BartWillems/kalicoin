FROM busybox:glibc

ENV ENVIRONMENT="development" \
    DATABASE_URI="postgres://user:pass@127.0.0.1:5432/kalicoin?sslmode=disable"

WORKDIR /home/kalicoin

COPY ./config /home/kalicoin/config
COPY ./migrations /home/kalicoin/migrations
COPY ./kalicoin /home/kalicoin/kalicoin

EXPOSE 8000

ENTRYPOINT [ "/home/kalicoin/kalicoin" ]