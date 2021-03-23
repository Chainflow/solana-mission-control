package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

type (
	// Telegram bot details struct
	Telegram struct {
		// BotToken is the token of your telegram bot
		BotToken string `mapstructure:"tg_bot_token"`
		// ChatID is the id of telegarm chat which will be used to get alerts
		ChatID int64 `mapstructure:"tg_chat_id"`
	}

	// SendGrid stores sendgrid API credentials
	SendGrid struct {
		// Token of sendgrid account
		Token string `mapstructure:"sendgrid_token"`
		// ToEmailAddress is the email to which all the alerts will be sent
		ReceiverEmailAddress string `mapstructure:"receiver_email_address"`
		// SendgridEmail is the email of sendgrid account which will be used to send mail alerts
		SendgridEmail string `mapstructure:"account_email"`
		// SendgridName is the name of sendgrid account which will be used to send mail alerts
		SendgridName string `mapstructure:"sendgrid_account_name"`
	}

	// Scraper defines the time intervals for multiple scrapers to fetch the data
	Scraper struct {
		// Rate is to call and get the data for specified targets on that particular time interval
		Rate string `mapstructure:"rate"`
		// ValidatorRate is to call and fetch the data from validatorStatus target on that time interval
		ValidatorRate string `mapstructure:"validator_rate"`
		// ContractRate is to call and fetch the data from smart contract realted targets on that time interval
		ContractRate string `mapstructure:"contract_rate"`
		// CommandsRate is to check the for telegram commands from telegram chat and returns the data
		CommandsRate string `mapstructure:"tg_commnads_rate"`
	}

	// Prometheus stores Prometheus details
	Prometheus struct {
		// Port on which influxdb is running
		Port string `mapstructure:"port"`
		// IP to connect to influxdb where it is running
		IP string `mapstructure:"ip"`
		// Database is the name of the influxdb database to store the data
		Database string `mapstructure:"database"`
		// Username is the name of the user of influxdb
		Username string `mapstructure:"username"`
		// Password of influxdb
		Password string `mapstructure:"password"`

		ListenAddress     string `mapstructure:"listen_address"`
		PrometheusAddress string `mapstructure:"prometheus_address"`
	}

	// Endpoints defines multiple API base-urls to fetch the data
	Endpoints struct {
		RPCEndpoint string `mapstructure:"rpc_endpoint"`
		NetworkRPC  string `mapstructure:"network_rpc"`
	}

	// ValDetails stores the validator meta details
	ValDetails struct {
		// ValidatorName is the moniker of your validator which will be used to display in alerts messages
		ValidatorName string `mapstructure:"validator_name"`
		PubKey        string `mapstructure:"pub_key"`
		VoteKey       string `mapstructure:"vote_key"`
	}

	// EnableAlerts struct which holds options to enalbe/disable alerts
	EnableAlerts struct {
		EnableTelegramAlerts bool `mapstructure:"enable_telegram_alerts"`
		EnableEmailAlerts    bool `mapstructure:"enable_email_alerts"`
	}

	// RegularStatusAlerts defines time-slots to receive validator status alerts
	RegularStatusAlerts struct {
		// AlertTimings is the array of time slots to send validator status alerts
		AlertTimings []string `mapstructure:"alert_timings"`
	}

	// AlerterPreferences stores individual alert settings to enable/disable particular alert
	AlerterPreferences struct {
		BalanceChangeAlerts        string `mapstructure:"balance_change_alerts"`
		AccountBalanceChangeAlerts string `mapstructure:"account_balance_change_alerts"`
		VotingPowerAlerts          string `mapstructure:"voting_power_alerts"`
		// ProposalAlerts             string `mapstructure:"proposal_alerts"`
		BlockDiffAlerts string `mapstructure:"block_diff_alerts"`
		// MissedBlockAlerts          string `mapstructure:"missed_block_alerts"`
		// NumPeersAlerts             string `mapstructure:"num_peers_alerts"`
		NodeSyncAlert   string `mapstructure:"node_sync_alert"`
		NodeStatusAlert string `mapstructure:"node_status_alert"`
		// EthLowBalanceAlert         string `mapstructure:"eth_low_balance_alert"`
		EpochDiffAlerts string `mapstructure:"epoch_diff_alrets"`
	}

	//  AlertingThreshold defines threshold condition for different alert-cases.
	// `Alerter` will send alerts if the condition reaches the threshold
	AlertingThreshold struct {
		// BlockDiffThreshold is to send alerts when the difference b/w network and validator
		// block height reaches the given threshold
		BlockDiffThreshold int64 `mapstructure:"block_diff_threshold"`
		// Alert when the validator identity balance is less than this amount of SOL
		AccountBalThreshold float64 `mapstructure:"account_bal_threshold"`

		EpochDiffThreshold int64 `mapstructure:"epoch_diff_threshold"`
	}

	// Config defines all the configurations required for the app
	Config struct {
		Endpoints           Endpoints           `mapstructure:"rpc_and_lcd_endpoints"`
		ValDetails          ValDetails          `mapstructure:"validator_details"`
		EnableAlerts        EnableAlerts        `mapstructure:"enable_alerts"`
		RegularStatusAlerts RegularStatusAlerts `mapstructure:"regular_status_alerts"`
		AlerterPreferences  AlerterPreferences  `mapstructure:"alerter_preferences"`
		AlertingThresholds  AlertingThreshold   `mapstructure:"alerting_threholds"`
		Scraper             Scraper             `mapstructure:"scraper"`
		Telegram            Telegram            `mapstructure:"telegram"`
		SendGrid            SendGrid            `mapstructure:"sendgrid"`
		Prometheus          Prometheus          `mapstructure:"prometheus"`
	}
)

// ReadFromFile to read config details using viper
func ReadFromFile() (*Config, error) {
	// usr, err := user.Current()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// configPath := path.Join(usr.HomeDir, `.matic-jagar/config/`)
	// log.Printf("Config Path : %s", configPath)

	v := viper.New()
	v.AddConfigPath(".")
	v.AddConfigPath("../")
	// v.AddConfigPath(configPath)
	v.SetConfigName("config")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("error while reading config.toml: %v", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("error unmarshaling config.toml to application config: %v", err)
	}

	if err := cfg.Validate(); err != nil {
		log.Fatalf("error occurred in config validation: %v", err)
	}

	return &cfg, nil
}

// Validate config struct
func (c *Config) Validate(e ...string) error {
	v := validator.New()
	if len(e) == 0 {
		return v.Struct(c)
	}
	return v.StructExcept(c, e...)
}
