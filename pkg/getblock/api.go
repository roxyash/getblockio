package getblock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/roxyash/getblock/internal/types"
)

func GetBlockNumber(apikey uuid.UUID) (types.ResponseBlockNumber, error) {
	client := &http.Client{}
	var data = strings.NewReader(`{"jsonrpc": "2.0",
									"method": "eth_blockNumber",
									"params": [],
									"id": "getblock.io"}`)
	req, err := http.NewRequest("POST", "https://eth.getblock.io/mainnet/", data)
	if err != nil {
		return types.ResponseBlockNumber{}, err
	}
	req.Header.Set("x-api-key", fmt.Sprintf("%v", apikey))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return types.ResponseBlockNumber{}, err
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return types.ResponseBlockNumber{}, err
	}

	responseBlockNumber := types.ResponseBlockNumber{}

	err = json.Unmarshal(bodyText, &responseBlockNumber)
	if err != nil {
		return types.ResponseBlockNumber{}, err
	}

	return responseBlockNumber, nil
}

func GetBlockInfoByNumber(apikey uuid.UUID, result string, hash bool) (types.ResponseBlockNumberInfo, error) {
	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf(`{"jsonrpc": "2.0",
									"method": "eth_getBlockByNumber",
									"params": ["%s", %v],
									"id": "getblock.io"}`, result, hash))
	req, err := http.NewRequest("POST", "https://eth.getblock.io/mainnet/", data)
	if err != nil {
		return types.ResponseBlockNumberInfo{}, err
	}
	req.Header.Set("x-api-key", fmt.Sprintf("%v",apikey))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return types.ResponseBlockNumberInfo{}, err
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return types.ResponseBlockNumberInfo{}, err
	}

	responseBlockNumberInfo := types.ResponseBlockNumberInfo{}

	err = json.Unmarshal(bodyText, &responseBlockNumberInfo)
	if err != nil {
		return types.ResponseBlockNumberInfo{}, err
	}

	return responseBlockNumberInfo, nil
}

func GetBalanceByMiner(apikey uuid.UUID, miner string) (types.ResponseBalanceInfo, error) {
	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf(`{"jsonrpc": "2.0",
									"method": "eth_getBalance",
									"params": ["%v", "latest"],
									"id": "getblock.io"}`, miner))
	req, err := http.NewRequest("POST", "https://eth.getblock.io/mainnet/", data)
	if err != nil {
		return types.ResponseBalanceInfo{}, err
	}
	req.Header.Set("x-api-key", fmt.Sprintf("%v", apikey))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return types.ResponseBalanceInfo{}, err
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return types.ResponseBalanceInfo{}, err
	}
	fmt.Printf("%s\n", bodyText)

	balanceInfo := types.ResponseBalanceInfo{}

	err = json.Unmarshal(bodyText, &balanceInfo)
	if err != nil {
		return types.ResponseBalanceInfo{}, err
	}

	return balanceInfo, nil
}



