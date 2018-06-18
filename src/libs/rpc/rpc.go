package rpc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Electra-project/electra-auth/src/helpers"
)

func query(method string, response interface{}) error {
	daemonURI := "http://127.0.0.1:5788"

	reqData := bytes.NewBuffer([]byte(`{"jsonrpc":"2.0","method":"` + method + `"}`))
	req, err := http.NewRequest("POST", daemonURI, reqData)
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("user", "pass")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		helpers.LogErr("Error: " + err.Error())

		return err
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		helpers.LogErr("Error: " + err.Error())

		return err
	}

	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		helpers.LogErr("Error: " + err.Error())

		return err
	}

	return nil
}
