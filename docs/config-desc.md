### Configure the following variables in `config.toml`
- **[rpc_and_lcd_endpoints]**
  - *rpc_endpoint*

      RPC end point is used to gather information like validator status,active stake, account balance, commission rate etc.
   - *network_rpc*

      NetworkRPC is used to gather information about network metrics like confirmed blocks, epoch information etc.

- **[validator_details]**

   - *validator_name*
   
      Validator name is the moniker of your validator which will be used to display in alert messages.

   - *pub_key*
  
      Public key of the validator, which will be used to gather information about validator metrics like commision,validator status etc...

   - *vote_key*
   
      Vote key of the validator, which will be used to get vote account details such as balance.

- **[enable_alerts]**

   - *enable_telegram_alerts*

      Configure **yes** if you wish to get telegram alerts otherwise make it **no**.

   - *enable_email_alerts*

      Configure **yes** if you wish to get email alerts otherwise make it **no**.

- **[alerter_preferences]**

   - *account_balance_change_alerts*

      Configure **yes** if you wish to get account balance change alerts otherwise make it **no**.

   - *block_diff_alerts*

      If you want to recieve alerts when there is a gap between your validator block height and network height then make it **yes** otherwise **no**.

   - *epoch_diff_alrets*

      If you want to recieve alerts when there is a gap between your validator epoch and network epoch then make it **yes** otherwise **no**.

   - *delegation_alerts*

      Configure **yes** if you wish to get alerts when there is a new delegation and your account balance changes, otherwise make it **no**.

   - *node_health_alert*
   
      If you want to recieve alerts when your validator node is down then make it as **yes** otherwise **no**.

   - *skip_rate_alerts*
     
      Configure **yes** if you wish to get alerts when validator skip rate exceeds network skip rate otherwise **no**.

- **[alerting_threholds]**

   - *block_diff_threshold*

      An Integer value to recieve block difference alerts, e.g. a value of 2 would alert you if your validator falls 2 or more blocks behind the network's current block height.

   - *epoch_diff_threshold*
       
      An integer value to recieve epoch difference alerts, e.g. a value of 5 would alert you if difference between your validator's epoch number and network's epoch is 5 or more.

   - *balance_change_threshold*

      An integer value to recieve account balance change alerts, e.g. if your account balance has dropped to given threshold value you will receive alerts.

- **[regular_status_alerts]**

   - *alert_timings*
   
      Array of timestamps for alerting about the validator health, i.e. whether it's voting or jailed. You can get alerts based on the time which can be configured.
     
- **[telegram]**
  - *tg_chat_id*

      Telegram chat ID to receive alerts to your telegram chat, required for Telegram alerting.
    
  - *tg_bot_token*

      Telegram bot token, required for Telegram alerting. The bot should be added to the chat and should have send message permission.
    
- **[Email]**

  - *email_address*

      E-mail address to receive mail notifications, required for e-mail alerting.
   
  - *sendgrid_token*

      Sendgrid mail service api token, required for e-mail alerting.

- **[prometheus]**

    - *prometheus_address*

      Prometheus address to export solana metrics and serve, by default listening address configured as (http://localhost:1234) in `config.toml`

    - *listen_address*
       
      Port in which prometheus server will run,and export metrics on this port, (ex: http://localhost:1234/metrics) shows all the metrics which are stored in prometheus database, by default it will run on 9090 port.
