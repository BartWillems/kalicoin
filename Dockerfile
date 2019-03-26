FROM busybox:glibc

ENV ENVIRONMENT="development" \
    DATABASE_URI="postgres://user:pass@127.0.0.1:5432/kalicoin?sslmode=disable"

WORKDIR /

COPY ./config /
COPY ./migrations /
COPY ./kalicoin /

EXPOSE 8000

ENTRYPOINT [ "/kalicoin" ]