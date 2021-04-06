
## Monitoring 

Monitoring is provided via three customized Grafana dashboards.  The dashboards provide consolidated, user-friendly, yet comprehensive views of a validator infrastucture's health.

Scroll past the dashboard descriptions to find a demo system link to see Solana Monitoring Tool in action 👇


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

**Note**: The above mentioned metrics will be calculated and displayed according to the validator address which will be configured in config.toml.      

### 2. System Monitoring Metrics

This view provides a comprehensive look at system performance metrics, expanding on the summary dashboard. Here's you'll find all the system metrics you'd expect to see in a comprehensive system monitoring tool.

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
