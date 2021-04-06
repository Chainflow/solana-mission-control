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
