package monitor

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/utils"
)

// TelegramAlerting will check for the commands from the configured telegram account
// If any commands are given in the tg account then Alerter will send the response back according to the input
func TelegramAlerting(cfg *config.Config) {
	if strings.ToUpper(strconv.FormatBool(cfg.EnableAlerts.EnableTelegramAlerts)) == "FALSE" {
		return
	}
	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.BotToken)
	if err != nil {
		log.Fatalf("Please configure telegram bot token %v:", err)
		return
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	msgToSend := ""

	for update := range updates {
		if update.Message == nil { // ignore if any non-Message Updates
			continue
		}

		if update.Message.Text == "/status" {
			msgToSend = GetStatus(cfg)
		} else if update.Message.Text == "/node" {
			msgToSend = NodeStatus(cfg)
		} else if update.Message.Text == "/balance" {
			msgToSend = GetAccountBal(cfg)
		} else if update.Message.Text == "/epoch" {
			msgToSend = GetEpochDetails(cfg)
		} else if update.Message.Text == "/vote_credits" {
			msgToSend = GetVoteCredits(cfg)
		} else if update.Message.Text == "/list" {
			msgToSend = GetHelp()
		} else {
			text := strings.Split(update.Message.Text, "")
			if len(text) != 0 {
				if text[0] == "/" {
					msgToSend = "Command not found do /list to know about available commands"
				} else {
					msgToSend = " "
				}
			}
		}

		log.Printf("[%s] %s", update.Message.From.UserName, msgToSend)

		if msgToSend != " " {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgToSend)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}

// GetHelp returns the msg to show for /help
func GetHelp() string {
	msg := "List of available commands\n /status - returns validator status, current block height " +
		"and network block height\n /node - return status of caught-up\n" +
		" /balance - returns the current balance of your account \n /epoch - returns current epoch of " +
		"network and validator\n /vote_credits - returns vote credits of current and" +
		"previous epochs \n /list - list out the available commands"

	return msg
}

// GetStatus returns the status messages for /status
func GetStatus(cfg *config.Config) string {
	var msg string

	status, err := GetValStatusFromDB(cfg)
	if err != nil {
		log.Printf("Error while getting validator status from db : %v", err)
	}

	msg = msg + fmt.Sprintf("Solana validator is currently %s\n", status)

	valHeight, err := GetEpochInfo(cfg, utils.Validator)
	if err != nil {
		log.Printf("Error while getting val block height res : %v", err)
	}
	msg = msg + fmt.Sprintf("Validator block height : %d\n", valHeight.Result.BlockHeight)

	networkHeight, err := GetEpochInfo(cfg, utils.Network)
	if err != nil {
		log.Printf("Error while getting network block height res : %v", err)
	}
	msg = msg + fmt.Sprintf("Network  block height : %d\n", networkHeight.Result.BlockHeight)

	return msg
}

// NodeStatus returns the node health wetaher it is up or down by giving /node
func NodeStatus(cfg *config.Config) string {
	var status string

	nodeHealth, err := GetNodeHealth(cfg) // Get solana node health
	if err != nil {
		log.Printf("Error while getting node health : %v", err)
	}

	if nodeHealth == 1 {
		status = fmt.Sprintf("- Your Solana validator node is %s \n", "UP")
	} else {
		status = fmt.Sprintf("- Your Solana validator node is %s \n", "DOWN")
	}

	return status
}

// GetAccountBal which resturns the account balance for the command /balance
func GetAccountBal(cfg *config.Config) string {
	var msg string

	res, err := GetBalance(cfg)
	if err != nil {
		log.Printf("Error while getting account balance : %v", err)
	}
	bal := float64(res.Result.Value) / math.Pow(10, 9)
	b := fmt.Sprintf("%.2f", bal)
	msg = fmt.Sprintf("Your account balance is %s SOL\n", b)

	return msg
}

// GetEpochDetails returns current epoch of validator and network for /epoch
func GetEpochDetails(cfg *config.Config) string {
	var msg string

	valEpoch, err := GetEpochInfo(cfg, utils.Validator)
	if err != nil {
		log.Printf("Error while getting val epoch info : %v", err)
	}

	msg = fmt.Sprintf("Current Epoch Info :: \n")

	msg = msg + fmt.Sprintf("Validator Epoch : %d\n", valEpoch.Result.Epoch)

	netEpoch, err := GetEpochInfo(cfg, utils.Network)
	if err != nil {
		log.Printf("Error while getting network epoch info : %v", err)
	}

	msg = msg + fmt.Sprintf("Network Epoch : %d\n", netEpoch.Result.Epoch)

	return msg
}

// GetVoteCredits returns credits for /vote_credits
func GetVoteCredits(cfg *config.Config) string {
	var msg string

	cCredits, pCredits, err := GetCredits(cfg)
	if err != nil {
		log.Printf("Error while getting vte credits from db : %v", err)
	}

	msg = fmt.Sprintf("Epoch vote credits ::\n Current epoch credits : %s\n Previous epoch credits: %s\n", cCredits, pCredits)

	return msg
}
