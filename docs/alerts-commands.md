## Alerting (Telegram and Email)
 A custom alerting module has been developed to alert on key validator health events. The module uses data from prometheus and triggers alerts based on user-configured thresholds.

 Here are the list of Alerts
 - Alert when node health is **DOWN**.
 - Alert when validator is in **DELINQUENT** state.
 - Alert when Block difference meets or exceedes **block_diff_threshold** which is user configured in *config.toml*.
 - Alert when Epoch difference reaches or exceedes **epoch_diff_threshold** which is user configured in *config.toml*.
 - Alert when account balance drops below **account_bal_threshold** which is user configured in *config.toml*.
 - Alert if validator skip rate exceedes network skip rate and difference of both exceedes **skip_rate_threshold** which is user configured in *config.toml* .

## Telegram Commands
These commands can be used to get quick information about your solana node.

Here is the list of available Telegram Commands.
  - **/list** - list out the available commands.
  - **/status** - status command returns validator status, current block height and network block height.
  - **/node** - return sync status.
  - **/balance** - returns the current balance of your account.
  - **/epoch** - returns current epoch of network and validator.
  - **/vote_credits** - returns vote credits of current and previous epochs.
  - **/rpc_status** - returns the status of validator rpc and network rpc i.e., running or not.
  - **/skip_rate** - returns the skip rate of validator and network.
  - **/block_production** - returns the recent block production details.
  - **/stop** - which panics the running code and also alerts will be stopped.
  
