# solana-prometheus

Solana prometheus monitoring tool provides a comprehensive set of metrics and alerts for solana validator node operators. We utilized the power of Grafana + Node exporter and extended the monitoring & alerting with a custom built go server. It also sends emergency alerts and calls to you based on your configuration.

## Install Prerequisites

- **Go 14.x+**
- **Grafana 7.x+**
- **Prometheus**
- **Node Exporter**


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

  - job_name: 'node_exporter'

    static_configs:
    - targets: ['localhost:9100']
    
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

```bash
$ sudo systemctl daemon-reload
$ sudo systemctl enable node_exporter.service
$ sudo systemctl start node_exporter.service
```
Click [here] to setup the customization ports for Grafana, Prometheus and Node Exporter. (./custom-port.md)
## Install and configure the Validator Mission Control

### Get the code
```bash
$ git clone https://github.com/PrathyushaLakkireddy/solana-prometheus
$ cd solana-prometheus
$ cp example.config.toml config.toml
```

### Configure the following variables in `config.toml`
- **[telegram]**
  - *tg_chat_id*

    Telegram chat ID to receive Telegram alerts, required for Telegram alerting.
    
  - *tg_bot_token*

    Telegram bot token, required for Telegram alerting. The bot should be added to the chat and should have send message permission.
- **[Email]**

  - *email_address*

    E-mail address to receive mail notifications, required for e-mail alerting.
- **[validator_details]**

   - *validator_name*
   
       Moniker of your validator,to get it displayed in alerts.

   - *pub_key*
  
      Public Node key of the validator, which will be used to get validator identity and other validator metrics like commision,validator status etc...

   - *vote_key*
   
      Vote key of validator, required for validator Identity.
    
- **[enable_alerts]**

   - *enable_telegram_alerts*

      Configure **yes** if you wish to get telegram alerts otherwise make it **no** .

   - *enable_email_alerts*

      Configure **yes** if you wish to get email alerts otherwise make it **no** .

- **[regular_status_alerts]**
   - *alert_timings*
   
      Array of timestamps for alerting about the validator health, i.e. whether it's voting or jailed. You can get alerts based on the time which can be configured.

- **[alerter_preferences]**

   - *account_balance_change_alerts*

       Configure **yes** if you wish to get account balance change alerts otherwise make it **no** .

   - *block_diff_alerts*

       If you want to recieve alerts when there is a gap between your validator block height and network height then make it **yes** otherwise **no**

   - *epoch_diff_alrets*

      If you want to recieve alerts when there is a gap between your validator epoch and network epoch then make it **yes** otherwise **no**

   - *delegation_alerts*

      Configure **yes** if you wish to get alerts about alters in account balance otherwise make it **no**

- **[alerting_threholds]**

   - *block_diff_threshold*

      An Integer value to recieve block difference alerts, e.g. a value of 2 would alert you if your validator falls 2 or more blocks behind the network's current block height.

    - *epoch_diff_threshold*
       
       An integer value to recieve epoch difference alerts, e.g. a value of 5 would alert you if difference between your validator's epoch number and network's epoch is 5 or more.

    - *account_bal_threshold*

       An integer value to recieve account balance change alerts, e.g. if your account balance has dropped to given threshold value you will receive alerts.

- **[prometheus]**

    - *prometheus_address*
       
       Address of your lcd client (ex: http://localhost:9090)

    - *listen_address*
       
       Port in which prometheus will run, by default it will run on 9090 port

After populating config.toml, check if you are running prometheus server in local

## Grafana Dashboards

Solana Monitoring Tool provides three dashboards

1. Validator Monitoring Metrics (These are the metrics which we have calculated and stored in prometheus)
2. System Metrics (These are the metrics related to the system configuration)
3. Summary (Which gives quick overview of validator and system metrics)

### 1. Validator monitoring metricss

The following list of metrics are displayed in this dashboard.

- **Validator Identity**

  Validator public Key:
    Node key of validator

  Validator vote key:
     Vote Key of the validator

- **Validator Information**

    Solana node version:
       Current version of the solana

    Solana node health:
        current health of the node

    IP Address:
        Gossip address of node
    
    Vote Account: 
       Shows information about whether the validator is voting or jailed
    
    Validator Active stake:
        Displays Validator current active stake
        
    Commision:
      Validator's vote account commision, percentage (0-100) of rewards payout owed to the vote account.

- **validator Health**

    Network Epoch:
       Displays network's epoch height 

    Validator Epoch:
       Displays validator's epoch height

    Epoch Difference:
       Difference between validator's and network's epoch

     Validator status:
        Shows validator's status whether the validator is voting or jailed

- **Validator Performance**

    Block Height - Network: 
       Displays the latest block height of a network
    
    Block Height - Validator:
       Displays the latest block height committed by the validator
    
    Height Difference:
       Displays Block Height Difference of network and validator's block height

    Solana Blocktime:
       Displays estimated production time of a confirmed block
    
    Vote Height - Network:
       Displays the latest vote height of network
    
    Vote Height - Validator:
       Displays the latest vote height committed by validator

    Vote Height Difference:
       Displays height difference of network and validator's vote height

    Account Balance:
       Displays Account Balance of validator in SOL's
    
    Solana current slot height:
        Displays Current slot height 
    
    Confirmed Blocktime - Network:
         Displays estimated production time of confirmed block of Network
    
    Confirmed Blocktime - Validator:
         Displays estimated production time of confirmed block committed by validator
        
    Block Time Difference:
        Difference between confirmation time of network and validator

    Current Epoch - vote credits:
        Displays total current epoch vote credits of validator vote account

    Previous Epoch - Vote credits:
        Displays total previous epoch vote credits of validator vote account
    
    Total valid slots:
       Displays number of leader valid slots per leader
    
    Total skipped slots:
       Displays number of leader skipped slots per leader

- **Validator Details**

   solana slot leader:
      Displays current slot leader address
   
   Transaction Count:
      Displays current Transaction count from the ledger
   
   Validator last voted:
      Displays Most recent slot voted on by this vote account
   
   Solana confirmed slot height:
      Displays current slot height
    
   Current Active Validators:
       Displays the number of current active validators, i.e validators who are voting 
    
   Delinquent validators:
      Displays the number of delinquent validators, i.e validators who are jailed.

   Confirmed epoch last slot:
      Displays current epoch's last slot

   Validator Root slot:
      Displays Root slot per validator

### 2. System Monitoring Metrics

These metrics are are collected by the node_exporter and displays all the metrics related to

- CPU
- Memory
- Disk
- Network traffic
- System processes
- Systemd
- Storage
- Hardware Misc
   
### 3. Summary Dashboard
This dashboard displays a quick information summary of validator details and system metrics. It includes following details.

- Validator identity (validator public key, validator vote key)
- Validator summary (Voting power, Validator sttus,Node Health,Block Height Difference) are the metrics being displayed from Validator details.
- CPU usage, RAM Usage, Memory usage and information about disk usage, Total RAM, CPU cores, server UPTime,CPU Basic, Memory Basic are the metrics being displayed from System details.

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


## Alerting (Telegram and Email)
 A custom alerting module has been developed to alert on key validator health events. The module uses data from influxdb and trigger alerts based on user-configured thresholds.

 - Alert when node health is **DOWN**
 - Alert when validator is in **DELINQUNET** state
 - Alert when Block difference meets **block_diff_threshold**
 - Alert when Epoch difference reaches to **epoch_diff_threshold**
 - Alert when there are alters in **Account Balance**
 - Alert when acount balance has dropped below to **account_bal_threshold**

## Telegram Commands
Telegram commands are used to query metric information on your telegram bot account

List of available Telegram Commands
  - **/status** - status command returns validator status, current block height and network block height
  - **/node** - return status of caught-up
  - **/balance** - returns the current balance of your account 
  - **/epoch** - returns current epoch of network and validator
  - **/vote_credits** - returns vote credits of current andprevious epochs 
  - **/rpc_status** - returns the status of validator rpc and network rpc i.e., running or not
  - **/stop** - which panics the running code and also alerts will be stopped
  - **/list** - list out the available commands




      
