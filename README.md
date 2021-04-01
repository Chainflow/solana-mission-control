# solana-prometheus

# Get the code
```bash
$ git clone https://github.com/PrathyushaLakkireddy/solana-prometheus
$ cd solana-prometheus
$ cp example.config.toml config.toml
```

# Configure the following variables in `config.toml`
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
