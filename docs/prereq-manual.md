### A - Install Grafana for Ubuntu
Download the latest .deb file and extract it by using the following commands

```sh
$ cd $HOME
$ sudo apt-get install -y adduser libfontconfig1
$ wget https://dl.grafana.com/oss/release/grafana_7.5.2_amd64.deb
$ sudo dpkg -i grafana_7.5.2_amd64.deb
```

Start the grafana server
```sh
$ sudo -S systemctl daemon-reload

$ sudo -S systemctl start grafana-server

Grafana will be running on port :3000 (ex:: https://localhost:3000)
```

### Install Prometheus

```sh
$ cd $HOME

$ wget https://github.com/prometheus/prometheus/releases/download/v2.22.1/prometheus-2.22.1.linux-amd64.tar.gz

$ tar -xvf prometheus-2.22.1.linux-amd64.tar.gz

$ sudo cp prometheus-2.22.1.linux-amd64/prometheus $GOBIN

$ sudo cp prometheus-2.22.1.linux-amd64/prometheus.yml $HOME
```
- Add the following in prometheus.yml using your editor of choices

```sh
 scrape_configs:

  - job_name: 'prometheus'

    static_configs:
    - targets: ['localhost:9090']
    
```

Setup Prometheus System service

```bash
sudo nano /lib/systemd/system/prometheus.service
```
- Copy-paste the following:
   
```sh
[Unit]
Description=Prometheus
After=network-online.target

[Service]
Type=simple
ExecStart=/home/ubuntu/go/bin/prometheus --config.file=/home/ubuntu/prometheus.yml
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```
- For the purpose of this guide it is assumed the `user` is `ubuntu`. If your user is   different please make the required changes above.
     
```sh 
$ sudo systemctl daemon-reload
$ sudo systemctl enable prometheus.service
$ sudo systemctl start prometheus.service
```

### Install node exporter

```sh
$ cd $HOME
$ curl -LO https://github.com/prometheus/node_exporter/releases/download/v0.18.1/node_exporter-0.18.1.linux-amd64.tar.gz
$ tar -xvf node_exporter-0.18.1.linux-amd64.tar.gz
$ sudo cp node_exporter-0.18.1.linux-amd64/node_exporter $GOBIN
```

Setup Node exporter service

```bash 
 sudo nano /lib/systemd/system/node_exporter.service
 ```

 Copy-paste the following:

 ```sh
 [Unit]
Description=Node_exporter
After=network-online.target

[Service]
Type=simple
ExecStart=/home/ubuntu/go/bin/node_exporter
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```
For the purpose of this guide it is assumed the `user` is `ubuntu`. If your user is different please make the required changes above.

- **Note**:  Do not forget to setup node exporter configuration in prometheus.yml file.
Copy paste the following in prometheus.yml.

- job_name: 'node_exporter'

    static_configs:
    - targets: [localhost:9100]
    
```bash
$ sudo systemctl daemon-reload
$ sudo systemctl enable node_exporter.service
$ sudo systemctl start node_exporter.service
```