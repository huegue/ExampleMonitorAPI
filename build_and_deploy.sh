#!/bin/bash

printf "Building Go binary...\n"
go build -o ./build/monitor_api ./api || {
    printf "Error: Failed to build Go binary\n"
    exit 1
}

printf "Deploying service...\n"

sudo mkdir -p /opt/monitor_api
sudo mv ./build/monitor_api /opt/monitor_api/monitor_api

sudo cp -r ./public /opt/monitor_api
sudo cp -r ./data /opt/monitor_api

sudo cp ./config/monitor_api.sh /opt/monitor_api/monitor_api.sh
sudo chmod +x /opt/monitor_api/monitor_api.sh

sudo cp ./config/monitor_api.service /etc/systemd/system/monitor_api.service

printf "Reloading systemd and starting monitor_api service...\n"
sudo systemctl daemon-reload
sudo systemctl enable monitor_api.service
sudo systemctl start monitor_api.service
sudo systemctl status monitor_api.service
