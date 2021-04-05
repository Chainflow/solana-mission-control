# solana-prometheus

## Monitoring 

Monitoring is provided via three customized Grafana dashboards.  The dashboards provide consolidated, user-friendly, yet comprehensive views of a validator infrastucture's health.

Scroll past the dashboard descriptions to find a demo system link to see Validator Mission Control in action 👇

### 1 - Summary Dashboard

This view provides a quick-look at overall validator and system health.

It shows you -

* Your validator's identifying information 
* Answers to these key validator health questions
* Am I voting?
* What's my voting power?
* Node Status 
* What is the Block Height difference between validator and network
* Critical system information, providing insight into memory, CPU and disk usage

### 2 - Validator Monitoring Dashboard

This view provides a comprehensive look at validator details and performance, expanding on the summary dashboard. It also includes proposal information.

It shows you -

* Validator identity 
* Validator information includes node version, node health, IP Address, validator voting or not, Active stake, commision.
* Validator Health includes validator and network epoch height, difference between validator and network epoch height, validator status
* Validator Performance which includes
    - Network and validator block height 
    - Block height difference
    - Block time 
    - Network and validator vote height  and difference between those heights
    - Account balance
    - Curent slot height 
    - Network and validator's confirmation time and difference between them.
    - Current and previous epoch vote credits
    - Total valid and skipped slots
* Validator Details which includes slot leader address, Total Transaction count, confirmed slot height, root slot, number of active and delinquent validators, last vote height of the validator

### 3 - System Monitoring Dashboard

This view provides a comprehensive look at system performance metrics, expanding on the summary dashboard. Here's you'll find all the system metrics you'd expect to see in a comprehensive system monitoring tool.

It shows you -

* CPU usage
* Memory usage
* Kernel performance
* Interrupts
* Network stack information (TCP/UDP)
* Network interface information
* Disk IOPS
* Disk space usage
* Metric velocity

## Alerting 

A custom-built alerting module complements the dashboards. The module provides configurable alerting that send warnings and alarms, when the validator systems and/or connectivity within the infrastructure experience issues.

The alerts are sent to a Telegram channel or email. Validators can update the code to send the alerts to any other communication channel you prefer.

Here's the full list of alerts -

* Alert when node health is DOWN
* Alert when validator is in delinquent state
* Alert when Block difference meets threshold which is configured in config
* Alert when Epoch difference reaches to **epoch_diff_threshold**
* Alert when there are alters in **Account Balance**
* Alert when acount balance has dropped below to **account_bal_threshold**


## Getting Started

[Setup Instructions](./INSTRUCTIONS.md)

