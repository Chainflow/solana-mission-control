package monitor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PrathyushaLakkireddy/solana-prometheus/alerter"
	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
)

// GetBalance returns the balance of the account
func GetBalance(cfg *config.Config) (types.Balance, error) {
	log.Println("Getting Account Balance...")
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.RPCEndpoint,
		Method:   http.MethodPost,
		Body: types.Payload{Jsonrpc: "2.0", Method: "getBalance", ID: 1, Params: []interface{}{
			cfg.ValDetails.PubKey, // should be base58 encoded to query data
		}},
	}

	var result types.Balance
	resp, err := HitHTTPTarget(ops)
	if err != nil {
		log.Printf("Error: %v", err)
		return result, err
	}

	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		log.Printf("Error: %v", err)
		return result, err
	}

	err = SendBalanceChangeAlert(result.Result.Value, cfg)
	if err != nil {
		log.Printf("Error while sending balance change alert : %v", err)
	}

	return result, nil
}

// GetIdentity returns the identity publickey for the current node on fail returns Error
func GetIdentity(cfg *config.Config) (types.Identity, error) {
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.RPCEndpoint,
		Method:   http.MethodPost,
		Body:     types.Payload{Jsonrpc: "2.0", Method: "getIdentity", ID: 1},
	}

	var result types.Identity
	resp, err := HitHTTPTarget(ops)
	if err != nil {
		log.Printf("Error: %v", err)
		return result, err
	}

	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		log.Printf("Error: %v", err)
		return result, err
	}

	return result, nil
}

// GetAccountBalFromDB get the account balance from DataBase
func GetAccountBalFromDB(cfg *config.Config) (string, error) {
	var result types.DBRes
	var bal string
	response, err := http.Get(fmt.Sprintf("%s/api/v1/query?query=solana_account_balance", cfg.Prometheus.PrometheusAddress))
	if err != nil {
		log.Printf("Error: %v", err)
		return bal, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(responseData, &result)
	if err != nil {
		log.Printf("Error: %v", err)
		return bal, err
	}
	if len(result.Data.Result) > 0 {
		bal = result.Data.Result[0].Metric.SolanaAccBalance
	}

	return bal, nil
}

// SendBalanceChangeAlert checks balance and DBbalance, If balance dropped to threshold,
// sends Alerts to the validator
func SendBalanceChangeAlert(currentBal int64, cfg *config.Config) error {
	prevBal, err := GetAccountBalFromDB(cfg)
	if err != nil {
		log.Printf("Error while getting bal from db : %v", err)
		return err
	}

	cBal := float64(currentBal)

	if strings.EqualFold(cfg.AlerterPreferences.AccountBalanceChangeAlerts, "yes") {
		if cBal < cfg.AlertingThresholds.AccountBalThreshold {
			err = alerter.SendTelegramAlert(fmt.Sprintf("Account Balance Alert: Your account balance has dropped below configured threshold, current balance is : %f", cBal), cfg)
			if err != nil {
				log.Printf("Error while sending account balance change alert : %v", err)
				return err
			}
		}
	}

	// Send delegation alerts
	if prevBal != "" {

		pBal, err := strconv.ParseFloat(prevBal, 64)
		if err != nil {
			log.Printf("Error while converting pBal to float64 : %v ", err)
			return err
		}

		if strings.EqualFold(cfg.AlerterPreferences.DelegationAlerts, "yes") {
			diff := cBal - pBal
			if diff > 0 {
				err = alerter.SendTelegramAlert(fmt.Sprintf("Delegation Alert: Your account balance has changed form %f to %f", pBal, cBal), cfg)
				if err != nil {
					log.Printf("Error while sending delegation alert : %v", err)
					return err
				}
			} else {
				err = alerter.SendTelegramAlert(fmt.Sprintf("Undelegation Alert: Your account balance has changed form %f to %f", pBal, cBal), cfg)
				if err != nil {
					log.Printf("Error while sending undelegation alert : %v", err)
					return err
				}
			}
		}
	}

	return nil
}
