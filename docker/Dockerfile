FROM golang:1.19

WORKDIR /src
COPY ./src .
RUN go mod download
COPY ./scripts /scripts
RUN chmod +x /scripts/*
EXPOSE 8080
ENTRYPOINT ["/scripts/entrypoint.sh"]
