#!/bin/bash

INSTALL_DIR="/opt/monitor_api"
BINARY="./monitor_api"

start() {
    cd "$INSTALL_DIR" || exit 1

    "$BINARY" --createdb >/dev/null 2>&1

    "$BINARY" --start >/dev/null 2>&1 &
}

stop() {
    PID=$(ps -ef | grep -i "$BINARY" | grep -v grep | awk '{print $2}' >/dev/null 2>&1)

    if [ -n "$PID" ]; then
        kill -s SIGKILL "$PID" >/dev/null 2>&1
        printf "Service stopped (PID: %s).\n" "$PID"
    else
        printf "Service is not running.\n"
    fi
}

case $1 in
start) start ;;
stop) stop ;;
*) printf "Usage: %s {start|stop}\n" "$0" ;;
esac
