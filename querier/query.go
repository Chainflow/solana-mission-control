package querier

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/solana-prometheus/config"
	"github.com/PrathyushaLakkireddy/solana-prometheus/types"
)

// GetAccountBalFromDB get the account balance from DataBase
func GetAccountBalFromDB(cfg *config.Config) (string, error) {
	var result types.DBRes
	var bal string
	response, err := http.Get(fmt.Sprintf("%s/api/v1/query?query=solana_account_balance", cfg.Prometheus.PrometheusAddress))
	if err != nil {
		log.Printf("Error while querying account bal from db: %v", err)
		return bal, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(responseData, &result)
	if err != nil {
		log.Printf("Error while unmarshelling account bal: %v", err)
		return bal, err
	}
	if len(result.Data.Result) > 0 {
		bal = result.Data.Result[0].Metric.SolanaAccBalance
	}

	log.Printf("account bal from db : %v", bal)

	return bal, nil
}

// AlertStatusCountFromPrometheus returns the AlertCount for validator voting alert
func AlertStatusCountFromPrometheus(cfg *config.Config) (string, error) {
	var result types.DBRes
	var count string
	response, err := http.Get(fmt.Sprintf("%s/api/v1/query?query=solana_val_alert_count", cfg.Prometheus.PrometheusAddress))
	if err != nil {
		log.Printf("Error: %v", err)
		return count, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(responseData, &result)
	if err != nil {
		log.Printf("Error: %v", err)
		return count, err
	}
	if len(result.Data.Result) > 0 {
		count = result.Data.Result[0].Metric.AlertCount
	}

	return count, nil
}

func GetValStatusFromDB(cfg *config.Config) (string, error) {
	var result types.DBRes
	var status string
	response, err := http.Get(fmt.Sprintf("%s/api/v1/query?query=solana_val_status", cfg.Prometheus.PrometheusAddress))
	if err != nil {
		log.Printf("Error: %v", err)
		return status, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(responseData, &result)
	if err != nil {
		log.Printf("Error: %v", err)
		return status, err
	}
	if len(result.Data.Result) > 0 {
		status = result.Data.Result[0].Metric.SolanaValStatus
	}

	return status, nil
}

// GetCredits returns the vote credits of previous and current epoch
func GetCredits(cfg *config.Config) (string, string, error) {
	var result types.DBRes
	var cCredits, pCredits string
	response, err := http.Get(fmt.Sprintf("%s/api/v1/query?query=solana_vote_credits", cfg.Prometheus.PrometheusAddress))
	if err != nil {
		log.Printf("Error: %v", err)
		return cCredits, pCredits, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error while reading vote credits from db : %v", err)
		return cCredits, pCredits, err
	}
	json.Unmarshal(responseData, &result)
	if err != nil {
		log.Printf("Error: %v", err)
		return cCredits, pCredits, err
	}
	if len(result.Data.Result) > 0 {
		cCredits = result.Data.Result[0].Metric.SolanaCurrentCredits
		pCredits = result.Data.Result[0].Metric.SolanaPreviousCredits
	}

	return cCredits, pCredits, nil
}
