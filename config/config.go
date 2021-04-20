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
	}

	// Prometheus stores Prometheus details
	Prometheus struct {
		// ListenAddress to export metrics on the given port
		ListenAddress string `mapstructure:"listen_address"`
		// PrometheusAddress to connect to prormetheus where it has running
		PrometheusAddress string `mapstructure:"prometheus_address"`
	}

	// Endpoints defines multiple API base-urls to fetch the data
	Endpoints struct {
		// RPCEndPoint is used to gather information about validator status,active stake, account balance, commission rate and etc.
		RPCEndpoint string `mapstructure:"rpc_endpoint"`
		// NetworkRPC is used to gather information about validator
		NetworkRPC string `mapstructure:"network_rpc"`
	}

	// ValDetails stores the validator meta details
	ValDetails struct {
		// ValidatorName is the moniker of your validator which will be used to display in alerts messages
		ValidatorName string `mapstructure:"validator_name"`
		// PubKey of validator as base-58 encoded string
		PubKey string `mapstructure:"pub_key"`
		// VoteKey of validator as base-58 encoded string
		VoteKey string `mapstructure:"vote_key"`
	}

	// EnableAlerts struct which holds options to enalbe/disable alerts
	EnableAlerts struct {
		// EnableTelegramAlerts which takes an option to enable/disable telegram alerts
		EnableTelegramAlerts bool `mapstructure:"enable_telegram_alerts"`
		// EnableTelegramAlerts which takes an option to enable/disable emial alerts
		EnableEmailAlerts bool `mapstructure:"enable_email_alerts"`
	}

	// RegularStatusAlerts defines time-slots to receive validator status alerts
	RegularStatusAlerts struct {
		// AlertTimings is the array of time slots to send validator status alerts at that particular timings
		AlertTimings []string `mapstructure:"alert_timings"`
	}

	// AlerterPreferences which holds individual alert settings which takes an option to  enable/disable particular alert
	AlerterPreferences struct {
		// DelegationAlerts which takes an option to disable/enable balance delegation alerts, on enable sends alert when current
		// account balance has dropped below from previous account balance.
		DelegationAlerts string `mapstructure:"delegation_alerts"`
		// AccountBalanceChangeAlerts which takes an option to disable/enable Account balance change alerts, on enable sends alert
		// when balance has dropped to balance threshold
		AccountBalanceChangeAlerts string `mapstructure:"account_balance_change_alerts"`
		// VotingPowerAlerts          string `mapstructure:"voting_power_alerts"`
		// BlockDiffAlerts which takes an option to enable/disable block height difference alerts, on enable sends alert
		// when difference meets or exceedes block difference threshold
		BlockDiffAlerts string `mapstructure:"block_diff_alerts"`
		// NodeHealthAlert which takes an option to  enable/disable node Health status alert, on enable sends alerts
		NodeHealthAlert string `mapstructure:"node_health_alert"`
		// NodeStatusAlert            string `mapstructure:"node_status_alert"`
		// EpochDiffAlerts which takes an option to enable/disable epoch difference alerts, on enable sends alerts if
		// difference reaches or exceedes epoch difference threshold
		EpochDiffAlerts string `mapstructure:"epoch_diff_alrets"`
	}

	//  AlertingThreshold defines threshold condition for different alert-cases.
	//`Alerter` will send alerts if the condition reaches the threshold
	AlertingThreshold struct {
		// BlockDiffThreshold is to send alerts when the difference b/w network and validator's
		// block height reaches or exceedes to block difference threshold
		BlockDiffThreshold int64 `mapstructure:"block_diff_threshold"`
		// AccountBalThreshold is to send Alert when the validator balance has dropped below to this threshold
		AccountBalThreshold float64 `mapstructure:"account_bal_threshold"`
		// EpochDiffThreahold option is to send alerts when the difference b/w network and validator's
		// epoch reaches or exceedes to epoch difference threshold
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

	// configPath := path.Join(usr.HomeDir, `.solana-tool/config/`)
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
