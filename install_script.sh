#!/bin/bash

set -e

cd $HOME

echo "----------- Installing grafana -----------"

sudo apt-get install -y adduser libfontconfig1

wget https://dl.grafana.com/oss/release/grafana_7.5.2_amd64.deb

sudo dpkg -i grafana_7.5.2_amd64.deb

echo "------ Starting grafana server using systemd --------"

sudo -S systemctl daemon-reload

sudo -S systemctl start grafana-server

cd $HOME

echo "----------- Installing Prometheus -----------"


# Make prometheus user
sudo adduser --no-create-home --disabled-login --shell /bin/false --gecos "Prometheus Monitoring User" prometheus

# Make directories and dummy files necessary for prometheus
sudo mkdir /etc/prometheus
sudo mkdir /var/lib/prometheus
sudo touch /etc/prometheus/prometheus.yml
sudo touch /etc/prometheus/prometheus.rules.yml

# Assign ownership of the files above to prometheus user
sudo chown -R prometheus:prometheus /etc/prometheus
sudo chown prometheus:prometheus /var/lib/prometheus

# Download prometheus and copy utilities to where they should be in the filesystem

wget https://github.com/prometheus/prometheus/releases/download/v2.22.1/prometheus-2.22.1.linux-amd64.tar.gz

tar -xvf prometheus-2.22.1.linux-amd64.tar.gz

sudo cp prometheus-2.22.1.linux-amd64/prometheus $GOBIN

sudo cp prometheus-2.22.1.linux-amd64/prometheus.yml $HOME


# Assign the ownership of the tools above to prometheus user
sudo chown -R prometheus:prometheus /etc/prometheus/consoles
sudo chown -R prometheus:prometheus /etc/prometheus/console_libraries
sudo chown prometheus:prometheus /usr/local/bin/prometheus
sudo chown prometheus:prometheus /usr/local/bin/promtool

# Populate configuration files
cat ./prometheus/prometheus.yml | sudo tee /etc/prometheus/prometheus.yml
cat ./prometheus/prometheus.rules.yml | sudo tee /etc/prometheus/prometheus.rules.yml
cat ./prometheus/prometheus.service | sudo tee /etc/systemd/system/prometheus.service

# systemd
sudo systemctl daemon-reload
sudo systemctl enable prometheus
sudo systemctl start prometheus

echo "-----------------Installing Node Exporter---------------------"

# Make node_exporter user
sudo adduser --no-create-home --disabled-login --shell /bin/false --gecos "Node Exporter User" node_exporter

# Download node_exporter and copy utilities to where they should be in the filesystem

curl -LO https://github.com/prometheus/node_exporter/releases/download/v0.18.1/node_exporter-0.18.1.linux-amd64.tar.gz

tar -xvf node_exporter-0.18.1.linux-amd64.tar.gz

sudo cp node_exporter-0.18.1.linux-amd64/node_exporter $GOBIN

sudo chown node_exporter:node_exporter /usr/local/bin/node_exporter

# systemd
cat ./node/node_exporter.service | sudo tee /etc/systemd/system/node_exporter.service

sudo systemctl daemon-reload
sudo systemctl enable node_exporter
sudo systemctl start node_exporter


echo "--------- Cloning cosmos-validator-mission-control -----------"

cd go/src/github.com

git clone https://github.com/PrathyushaLakkireddy/solana-prometheus

cd solana-prometheus

cp example.config.toml config.toml

echo "------ Building and running the code --------"

go build && ./solana-prometheus