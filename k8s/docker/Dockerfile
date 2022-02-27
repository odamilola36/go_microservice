FROM golang

WORKDIR /app

COPY authsrv /app/authsrv

COPY apisrv /app/apisrv

EXPOSE 9000

EXPOSE 9001

CMD ["/app/authsrv", "/app/apisrv"]