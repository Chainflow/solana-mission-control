## Alerting (Telegram and Email)
 A custom alerting module has been developed to alert on key validator health events. The module uses data from prometheus and trigger alerts based on user-configured thresholds.

 Here are the list of Alerts
 - Alert when node health is **DOWN**
 - Alert when validator is in **DELINQUNET** state
 - Alert when Block difference meets or exceedes **block_diff_threshold** which is user configured in *config.toml*
 - Alert when Epoch difference reaches or exceedes **epoch_diff_threshold** which is user configured in *config.toml*
 - Alert when Account balance has dropped from previous **Account Balance** to current **Account_Balance** in SOL's.
 - Alert when acount balance has dropped below to **account_bal_threshold** which is user configured in *config.toml*

## Telegram Commands
Telegram commands will be used to get a quick information about your solana node. Based on the commands you will get alerts to your account.

Here is the list of available Telegram Commands.
  - **/status** - status command returns validator status, current block height and network block height
  - **/node** - return status of caught-up
  - **/balance** - returns the current balance of your account 
  - **/epoch** - returns current epoch of network and validator
  - **/vote_credits** - returns vote credits of current andprevious epochs 
  - **/rpc_status** - returns the status of validator rpc and network rpc i.e., running or not
  - **/skip_rate** - returns the skip rate of validator and network
  - **/block_production** - returns the recent block production details
  - **/stop** - which panics the running code and also alerts will be stopped
  - **/list** - list out the available commands
