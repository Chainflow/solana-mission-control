#!/bin/bash

set -e


cd $HOME

echo "------ checking for go, if it's not installed then it will be installed here -----"

command_exists () {
    type "$1" &> /dev/null ;
}

if command_exists go ; then
    echo "Golang is already installed"
else
  echo "------- Install dependencies -------"
  sudo apt update
  sudo apt install build-essential jq -y

  wget https://dl.google.com/go/go1.15.5.linux-amd64.tar.gz
  tar -xvf go1.15.5.linux-amd64.tar.gz
  sudo mv go /usr/local

  echo "------ Update bashrc ---------------"
  export GOPATH=$HOME/go
  export GOROOT=/usr/local/go
  export GOBIN=$GOPATH/bin
  export PATH=$PATH:/usr/local/go/bin:$GOBIN
  echo "" >> ~/.bashrc
  echo 'export GOPATH=$HOME/go' >> ~/.bashrc
  echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
  echo 'export GOBIN=$GOPATH/bin' >> ~/.bashrc
  echo 'export PATH=$PATH:/usr/local/go/bin:$GOBIN' >> ~/.bashrc

  source ~/.bashrc

  mkdir -p "$GOBIN"
fi

cd $HOME

echo "----------- Installing grafana -----------"

sudo apt-get install -y adduser libfontconfig1

wget https://dl.grafana.com/oss/release/grafana_7.5.2_amd64.deb

sudo dpkg -i grafana_7.5.2_amd64.deb

echo "------ Starting grafana server using systemd --------"

sudo -S systemctl daemon-reload

sudo -S systemctl start grafana-server

cd $HOME

echo "----------- Installing prometheus -----------"

wget https://github.com/prometheus/prometheus/releases/download/v2.22.1/prometheus-2.22.1.linux-amd64.tar.gz
tar -xvf prometheus-2.22.1.linux-amd64.tar.gz

tar -xvf prometheus-2.22.1.linux-amd64.tar.gz

cp prometheus-2.22.1.linux-amd64/prometheus $GOBIN

cp prometheus-2.22.1.linux-amd64/prometheus.yml $HOME


echo "------- Edit prometheus.yml --------------"

echo "
  
  - job_name: 'prometheus'
    static_configs:
    - targets: ['localhost:9090']
  - job_name: 'node_exporter'
    static_configs:
    - targets: [localhost:9100]" >> "$HOME/prometheus.yml"

if [ -z "${SENTRY1}" ];
then
   echo "-----Sentry-1 IP is empty, so not writing into prometheus.yml--------"
else 
  echo "
  - job_name: 'sentry-1'
    static_configs:
    - targets: ['$SENTRY1:26660']" >> "$HOME/prometheus.yml"
fi

if [ -z "${SENTRY2}" ];
then
  echo "-----Sentry-2 IP is empty, so not writing into prometheus.yml--------"
else
  echo "
  - job_name: 'sentry-2'
    static_configs:
    - targets: ['$SENTRY2:26660']" >> "$HOME/prometheus.yml"
fi

echo "------- Setup prometheus system service -------"

echo "[Unit]
Description=Prometheus
After=network-online.target
[Service]
Type=simple
ExecStart=$HOME/go/bin/prometheus --config.file=$HOME/prometheus.yml
Restart=always
RestartSec=3
LimitNOFILE=4096
[Install]
WantedBy=multi-user.target" | sudo tee "/lib/systemd/system/prometheus.service"

echo "------ Start prometheus -----------"

sudo systemctl daemon-reload
sudo systemctl enable prometheus.service
sudo systemctl start prometheus.service

echo "-------- Installing node exporter -----------"

cd $HOME

curl -LO https://github.com/prometheus/node_exporter/releases/download/v0.18.1/node_exporter-0.18.1.linux-amd64.tar.gz

tar -xvf node_exporter-0.18.1.linux-amd64.tar.gz

cp node_exporter-0.18.1.linux-amd64/node_exporter $GOBIN

echo "---------- Setup Prometheus Node exporter service -----------"

echo "[Unit]
Description=Node_exporter
After=network-online.target
[Service]
Type=simple
ExecStart=$HOME/go/bin/node_exporter
Restart=always
RestartSec=3
LimitNOFILE=4096
[Install]
WantedBy=multi-user.target" | sudo tee "/lib/systemd/system/node_exporter.service"

echo "----------- Start node exporter ------------"

sudo systemctl daemon-reload

sudo systemctl enable node_exporter.service

sudo systemctl start node_exporter.service

echo "---- Cleaning .dep .tar.gz files of grafana, prometheus, influxdb and node exporter --------"

rm influxdb_1.8.3_amd64.deb grafana_7.3.1_amd64.deb node_exporter-0.18.1.linux-amd64.tar.gz prometheus-2.22.1.linux-amd64.tar.gz

echo "** Done with prerequisite installtion **"
