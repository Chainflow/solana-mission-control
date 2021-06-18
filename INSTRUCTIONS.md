# solana-prometheus

Solana prometheus monitoring tool provides a comprehensive set of metrics and alerts for solana validator node operators. We utilized the power of Grafana + Node exporter and extended the monitoring & alerting with a custom built go server.

## Install Prerequisites

- **Go 14.x+**
- **Grafana 7.x+**
- **Prometheus**
- **Node Exporter**
- **Solana Client Binary**

### Prerequisite Installation

There are two ways of installing the prerequisites:-

   1. Installation script
   2. Manual installation

Either of the two methods can be used to install the required prerequisites. It is not necessary to do both.

**1) Installation script**

   - Script downloads and installs grafana, prometheus and node exporter and starts the respective servers.
   - It also downloads go if it's not already installed.
   - The script takes env variables and writes them to `config.toml` file.
   
   - You can find the script [here](./scripts/install_script.sh)
   - Execute the script using the following command:

   ```sh
   curl -s -L https://github.com/Chainflow/solana-mission-control/tree/main/scripts/install_script.sh | bash
   ```
   Source your `.bashrc` after executing the script

   ```sh
   source ~/.bashrc
   ```
   **Note**: This script installs the prerequisites and enables them to run on their default ports ie. `Grafana` by default runs on port 3000, `prometheus` by default runs on port 9090 and `Node Exporter` by default runs on port 9100. If you want to change the defaults ports please follow these [Instructions](./docs/custom-port.md)

   You can view the logs by executing the following commands:
   ```bash
   journalctl -u grafana-server -f

   journalctl -u prometheus.service -f

   journalctl -u node_exporter.service -f
   ```

**2) Manual installation**

To manually install the prerequisites please follow this [guide](./docs/prereq-manual.md).

 
## Install and configure the Solana Monitoring Tool

There are two ways of installing the tool:-

1. Installation script
2. Manual installation
Either of the two methods can be used to install the tool. It is not necessary to do both.

**1) Installation script**

  - It clones and sets up the monitoring tool as a system service.
  - Please export the following env variables first as they will be used to initialize the `config.toml` file for the tool.
  ```sh
  cd $HOME
  export RPC_ENDPOINT="<validator-endpoint>" # Ex - export RPC_ENDPOINT="https://api.xxxxxxxxxxxxxxxxxxxx.com"
  export NETWORK_RPC="<network-endpoint>" # Ex - export NETWORK_RPC="https://api.xxxxxxxxxxxxxxxxxxxx.com"
  export VALIDATOR_NAME="<moniker>" # Your validator name
  export PUB_KEY="<node-Public-key>"  # Ex - export PUB_KEY="valmmK7i1AxXeiTtQgQZhQNiXYU84ULeaYF1EH1pa"
  export VOTE_KEY="<vote-key>" # Ex - export VOTE_KEY="2oxQJ1qpgUZU9JU84BHaoM1GzHkYfRDgDQY9dpH5mghh"
  export TELEGRAM_CHAT_ID=<id> # Ex - export TELEGRAM_CHAT_ID=22828812
  export TELEGRAM_BOT_TOKEN="<token>" # Ex - TELEGRAM_BOT_TOKEN="1117273891:AAEtr3ZU5x4JRj5YSF5LBeu1fPF0T4xj-UI"
```
- **Note**: if you don't want telegram notifications you can skip exporting `TELEGRAM_CHAT_ID` and `TELEGRAM_BOT_TOKEN` but the rest are mandatory.
- You can find the tool installation script [here](./scripts/tool_installation.sh)
- Run the script using the following command

```sh
   curl -s -L https://github.com/Chainflow/solana-mission-control/tree/main/scripts/scripts/tool_installation.sh | bash
```
You can check the logs of tool using:
```sh
   journalctl -u solana_prometheus.service
```
### 2) Manual installation

```bash
$ git clone https://github.com/Chainflow/solana-mission-control
$ cd solana-mission-control
$ cp example.config.toml config.toml
```
Edit the `config.toml` with your changes. Information on all the fields in `config.toml` can be found [here](./docs/config-desc.md)

Note : Before running this monitoring binary, you need to add below configuration to prometheus.yml, which you have configured while installing prometheus.

```
- Add the following in prometheus.yml using your editor of choices

```sh
 scrape_configs:

  - job_name: 'Solana'

    static_configs:
    - targets: ['localhost:1234']

```

Then, restart the prometheus serivce

```sh 
$ sudo systemctl daemon-reload
$ sudo systemctl start prometheus.service
```

- Now you can build and run the monitoring binary

```sh
   $ go build -o solana-prometheus && ./solana-prometheus
```

Installation of the tool is completed lets configure the Grafana dashboards.

### Grafana Dashboards

The repo provides three dashboards

1. Validator Monitoring Metrics - Displays the validator metrics which are calculated and stored in prometheus.
2. System Metrics - Displays the metrics related to your validator server on which this tool is hosted on.
3. Summary - Displays a quick overview of validator monitoring metrics and system metrics.

Information of all the dashboards can be found [here](./docs/dashboard-desc.md).

## How to import these dashboards in your Grafana installation

### 1. Login to your Grafana dashboard
- Open your web browser and go to http://<your_ip>:3000/. `3000` is the default HTTP port that Grafana listens to if you haven’t configured a different port.
- If you are a first time user type `admin` for the username and password in the login page.
- You can change the password after login.

### 2. Create Datasource

- Before importing the dashboards you have to create datasources of `Prometheus`.

- To create datasoruces go to configuration and select Data Sources.

- Click on Add data source and select `Prometheus` from Time series databases section.

- Replace the URL with http://localhost:9090. 

- Click on Save & Test . Now you have a working Datasource of Prometheus.

### 3. Import the dashboards

- To import the json file of the **validator monitoring metrics** click the *plus* button present on left hand side of the dashboard. Click on import and load the validator_monitoring_metrics.json present in the grafana_template folder. 

- Select the datasources and click on import.

- To import **system monitoring metrics** click the *plus* button present on left hand side of the dashboard. Click on import and load the system_monitoring_metrics.json present in the grafana_template folder.

- While creating this dashboard if you face any issues at valueset, change it to empty and then click on import by selecting the datasources.

- To import **summary**, click the *plus* button present on left hand side of the dashboard. Click on import and load the summary.json present in the grafana_template folder.

- *For more info about grafana dashboard imports you can refer https://grafana.com/docs/grafana/latest/reference/export_import/*






      
