FROM scratch

ENV ENVIRONMENT="development" \
    DATABASE_URI="postgres://user:pass@127.0.0.1:5432/kalicoin?sslmode=disable"


ADD ./kalicoin /kalicoin

EXPOSE 8000

ENTRYPOINT [ "/kalicoin" ]