FROM golang:1.23 AS builder

WORKDIR /app_build

COPY go.mod go.sum ./
RUN go mod download

COPY api ./api
RUN go build -o monitor_api ./api

FROM ubuntu:22.04

WORKDIR /opt/monitor_api

COPY --from=builder /app_build/monitor_api ./

COPY public ./public
COPY data ./data

COPY docker_entrypoint.sh ./
RUN chmod +x ./docker_entrypoint.sh

EXPOSE 8090

VOLUME /opt/monitor_api/data

CMD ["./docker_entrypoint.sh"]




