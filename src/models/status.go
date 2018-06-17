package models

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Electra-project/electra-auth/src/helpers"
)

// Status model.
type Status struct {
	Result StatusResult `json:"result"`
	Error  string       `json:"error"`
	ID     string       `json:"id"`
}

// StatusResult model.
type StatusResult struct {
	Version         string                 `json:"version"`
	ProtocolVersion int64                  `json:"protocolversion"`
	WalletVersion   float64                `json:"walletversion"`
	Balance         float64                `json:"balance"`
	NewMint         float64                `json:"newmint"`
	Stake           float64                `json:"stake"`
	Blocks          int64                  `json:"blocks"`
	TimeOffset      int64                  `json:"timeoffset"`
	MoneySupply     float64                `json:"moneysupply"`
	Connections     int64                  `json:"connections"`
	Proxy           string                 `json:"proxy"`
	IP              string                 `json:"ip"`
	Difficulty      StatusResultDifficulty `json:"difficulty"`
	Testnet         bool                   `json:"testnet"`
	KeyPoolOldest   int64                  `json:"keypoololdest"`
	KeyPoolSize     int64                  `json:"keypoolsize"`
	PayTxFee        float64                `json:"paytxfee"`
	MinInput        float64                `json:"mininput"`
	UnlockedUntil   int64                  `json:"unlocked_until"`
	Errors          string                 `json:"errors"`
}

// StatusResultDifficulty model.
type StatusResultDifficulty struct {
	ProofOfWork  float64 `json:"proof-of-work"`
	ProofOfStake float64 `json:"proof-of-stake"`
}

// Get gets the daemon status info.
func (s Status) Get() (*Status, error) {
	daemonURI := "http://127.0.0.1:5788"

	reqData := bytes.NewBuffer([]byte(`{"jsonrpc":"2.0","method":"getinfo"}`))
	req, err := http.NewRequest("POST", daemonURI, reqData)
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("user", "pass")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		helpers.LogErr("Error: " + err.Error())

		return nil, err
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		helpers.LogErr("Error: " + err.Error())

		return nil, err
	}

	var status *Status
	err = json.Unmarshal(bodyBytes, &status)
	if err != nil {
		helpers.LogErr("Error: " + err.Error())

		return nil, err
	}

	return status, nil
}
