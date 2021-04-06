# solana-prometheus

Solana prometheus monitoring tool provides a comprehensive set of metrics and alerts for solana validator node operators. We utilized the power of Grafana + Node exporter and extended the monitoring & alerting with a custom built go server. It also sends emergency alerts and calls to you based on your configuration.

## Install Prerequisites

- **Go 14.x+**
- **Grafana 7.x+**
- **Prometheus**
- **Node Exporter**

To install the prerequistes please follow this [guide](./docs/prereq-manual.md).

 
## Install and configure the Solana Monitoring Tool

### Get the code
```bash
$ git clone https://github.com/PrathyushaLakkireddy/solana-prometheus
$ cd solana-prometheus
$ cp example.config.toml config.toml
```
Edit the `config.toml` with your changes. Information on all the fields in `config.toml` can be found [here](./docs/config-desc.md)

Installation of the tool is completed lets configure the Grafana dashboards.

### Grafana Dashboards

The repo provides five dashboards

1. Validator Monitoring Metrics - Displays the validator metrics which are calculated and stored in prometheus.
2. System Metrics - Displays the metrics related to your validator server on which this tool is hosted on.
3. Summary - Displays a quick overview of heimdall, bor and system metrics.

Information on all the dashboards can be found [here](./docs/dashboard-desc.md).

## How to import these dashboards in your Grafana installation

### 1. Login to your Grafana dashboard
- Open your web browser and go to http://<your_ip>:3000/. `3000` is the default HTTP port that Grafana listens to if you haven’t configured a different port.
- If you are a first time user type `admin` for the username and password in the login page.
- You can change the password after login.

### 2. Create Datasource

- Before importing the dashboards you have to create datasources of Prometheuss.
- To create datasoruces go to configuration and select Data Sources.
- After that you can find Add data source, select Prometheus from Time series databases section.

### 3. Import the dashboards
- To import the json file of the **validator monitoring metrics** click the *plus* button present on left hand side of the dashboard. Click on import and load the validator_monitoring_metrics.json present in the grafana_template folder. 

- Select the datasources and click on import.

- To import **system monitoring metrics** click the *plus* button present on left hand side of the dashboard. Click on import and load the system_monitoring_metrics.json present in the grafana_template folder.

- While creating this dashboard if you face any issues at valueset, change it to empty and then click on import by selecting the datasources.

- To import **summary**, click the *plus* button present on left hand side of the dashboard. Click on import and load the summary.json present in the grafana_template folder.

- *For more info about grafana dashboard imports you can refer https://grafana.com/docs/grafana/latest/reference/export_import/*






      
