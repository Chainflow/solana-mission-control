
## Monitoring 

Monitoring is provided via three customized Grafana dashboards.  The dashboards provide consolidated, user-friendly, yet comprehensive views of a validator infrastucture's health.

Scroll past the dashboard descriptions to find a demo system link to see Solana Monitoring Tool in action 👇


### 1. Validator monitoring metricss

The following list of metrics are displayed in this dashboard.

- **Validator Identity**

  *Pub Key*: Node public key of the validator.

  *Vote key*: Vote Key of the validator.

- **Validator Information**

    *Node Version*: Current version of the solana node.

    *Node health*: Current health of the solana node, if it is running fine, then it marked as **UP** or else marked as **DOWN**.

    *Validator Active stake*: Displays the stake, delegated to the vote account and active in the current epoch.

    *IP Address*: Gossip network address for the node.
    
    *Vote Account*: Shows information about whether the validator is voting or jailed, If the validator is voting, it marked as **YES** or else **NO**.
        
    *Commision*: Validator's vote account commision, percentage (0-100) of rewards payout owed to the vote account.

- **Validator Health**

    *Network Epoch*: Displays network's epoch height .

    *Validator Epoch*: Displays validator's epoch height.

    *Epoch Difference*: Difference between validator's and network's epoch height.

    *Vote Height - Network*: Displays the latest vote height of network.
    
    *Vote Height - Validator*: Displays the latest vote height committed by validator.

    *Vote Height Difference*: Displays height difference of network and validator's vote height.

    *Block Height - Network*: Displays the latest block height of a network.
    
    *Block Height - Validator*: Displays the latest block height committed by the validator.
    
    *Block Height Difference*: Displays Block Height Difference of network and validator.

    *Voting*: Shows validator's status whether the validator is voting or jailed, If the validator is voting then it displays as **VOTING** or else **JAILED**.

    *Skip Rate - Validator*: Displays skip rate of validator.

    *Skip Rate - Network*: Displays skip Rate of network.
    
    *Skip Rate - Difference*: Displays difference between validator and network skip rate.

    *Vote Height - Network*: Displays the latest vote height of network.
    
    *Vote Height - Validator*: Displays the latest vote height committed by validator.

    *Vote Height Difference*: Displays height difference of network and validator's vote height.

- **Validator Performance**

    *Current Epoch - Vote credits*: Displays vote credits of vote account for current epoch.

    *Previous Epoch - Vote credits*: Displays vote credits of validator for previous epoch.
    
    *Identity Account Balance*: Displays identity account balance.

    *Confirmed Blocktime - Network*: Displays estimated production time of confirmed block of a network.
    
    *Confirmed Blocktime - Validator*: Displays estimated production time of confirmed block committed by validator.
        
    *Confirmed Time Difference*: Difference between confirmation time of validator and network.

    *Total Transaction Count*: Displays total transaction count from the current ledger.
    
    *Confirmed Epoch Last Slot - Validator*: Displays confirmed epoch of validator in last slot.

    *Confirmed Epoch Last Slot - Network*: Displays confirmed epoch of network in last slot.

    *Vote Account Balance*: Displays current vote account balance.

- **Recent Block Production - Current Epoch**

    *Leader Slots - Validator*:Displays the no.of leader slots of a validator in current epoch.

    *Total Slots - Current Epoch*:Displays total slots in current epoch.

    *Blocks Produced - Validator*: Displays produced blocks of a validator in current epoch.

    *Total Blocks Produced - Current Epoch*: Displays total blocks produced in current epoch.

    *Skipped Slots - Validator*: Displays Skipped slots of a validator in current epoch.

    *Total Skipped Slots - Current Epoch* - Displays Total skipped slots in current epoch.

- **Extra Information**

    *Solana Slot Leader*: Display slot leader for current slot.

    *Delinquent validators*: Displays the number of delinquent validators, i.e validators who are not active. 

    *Validator Root Slot*: Displays root slot of validator.

    *Last voted*: Displays recent slot voted by vote account.

    *Current Active Validators*: Displays the number of current active validators, i.e validators who are voting.

    *Confirmed Slot Height*: Displays confirmed slot height.

**Note**: The above mentioned metrics will be calculated and displayed according to the addresses which will be configured in config.toml.      

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
- Validator summary which includes  
         Voting Account, Validator status,Node status,
         Block Height - Network, Block Height - Validator, Block Height Difference 
         Validator Commision, Active stake, Current Epoch Vote Credits
- Server Uptime, CPU Busy, RAM Used, CPU, Memory Stake, Network Traffic, Disk Space Used, Disk IOps, I/O usage Read/Write, I/O Usage Times are the metrics being displayed from System details.
