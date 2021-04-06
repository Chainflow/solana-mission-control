## Default Port Customization

### Grafana:-

To change the default port of your Grafana server please edit the following files: `/usr/share/grafana/conf/defaults.ini` and `/etc/grafana/grafana.ini`. Search for `3000` and replace it with your custom port. Restart the systemd service after making your changes.

```sh
sudo systemctl restart grafana-server
```

### Prometheus:-

To change the default port of `Prometheus` process add this flag `--web.listen-address="0.0.0.0:<port>"` in `/lib/systemd/system/prometheus.service` file. Ex -

```sh
[Service]
Type=simple
ExecStart=$HOME/go/bin/prometheus --config.file=$HOME/prometheus.yml --web.listen-address="0.0.0.0:5000"
Restart=always
RestartSec=3
LimitNOFILE=4096
```

### Node Exporter:-

To change the default port of `Node_exporter` process add this flag `--web.listen-address="0.0.0.0:<port>"` in `/lib/systemd/system/node_exporter.service` file. Ex -

```sh
[Service]
Type=simple
ExecStart=$HOME/go/bin/node_exporter --web.listen-address="0.0.0.0:5000"
Restart=always
RestartSec=3
LimitNOFILE=4096
```
For the changes to take effect reload the file and restart the process.

```bash
sudo systemctl daemon-reload
sudo systemctl restart node_exporter.service
```
- **Note**: Please make sure to edit the `prometheus.yml` with your custom port for `node_exporter` scrape job.

