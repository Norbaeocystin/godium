package godium

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// get pools data via raydium api
func GetPoolsViaApi() (PoolsV3Data, error) {
	var poolsData PoolsV3Data
	client := http.Client{}
	request, err := http.NewRequest("GET", POOLS_V3_API, nil)
	if err != nil {
		return poolsData, err
	}
	request.Header.Set("Content-Type", "application/json")
	// Make request
	response, err := client.Do(request)
	if err != nil {
		return poolsData, err
	}
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return poolsData, err
	}
	defer response.Body.Close()
	err = json.Unmarshal(bodyBytes, &poolsData)
	return poolsData, err
}

func GetTokens() (TokenList, error) {
	var tokenData TokenList
	client := http.Client{}
	request, err := http.NewRequest("GET", TOKENS_API, nil)
	if err != nil {
		return tokenData, err
	}
	request.Header.Set("Content-Type", "application/json")
	// Make request
	response, err := client.Do(request)
	if err != nil {
		return tokenData, err
	}
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return tokenData, err
	}
	defer response.Body.Close()
	err = json.Unmarshal(bodyBytes, &tokenData)
	return tokenData, err
}

type PoolsV3Data struct {
	Data []struct {
		ID            string `json:"id"`
		MintA         string `json:"mintA"`
		MintB         string `json:"mintB"`
		VaultA        string `json:"vaultA"`
		VaultB        string `json:"vaultB"`
		MintDecimalsA int    `json:"mintDecimalsA"`
		MintDecimalsB int    `json:"mintDecimalsB"`
		AmmConfig     struct {
			ID              string `json:"id"`
			Index           int    `json:"index"`
			ProtocolFeeRate int    `json:"protocolFeeRate"`
			TradeFeeRate    int    `json:"tradeFeeRate"`
			TickSpacing     int    `json:"tickSpacing"`
			FundFeeRate     int    `json:"fundFeeRate"`
			FundOwner       string `json:"fundOwner"`
			Description     string `json:"description"`
		} `json:"ammConfig"`
		Tvl float64 `json:"tvl"`
		Day struct {
			Volume    float64 `json:"volume"`
			VolumeFee float64 `json:"volumeFee"`
			FeeA      float64 `json:"feeA"`
			FeeB      float64 `json:"feeB"`
			FeeApr    float64 `json:"feeApr"`
			RewardApr struct {
				A float64 `json:"A"`
				B float64 `json:"B"`
				C float64 `json:"C"`
			} `json:"rewardApr"`
			Apr      float64 `json:"apr"`
			PriceMin float64 `json:"priceMin"`
			PriceMax float64 `json:"priceMax"`
		} `json:"day"`
		Week struct {
			Volume    float64 `json:"volume"`
			VolumeFee float64 `json:"volumeFee"`
			FeeA      float64 `json:"feeA"`
			FeeB      float64 `json:"feeB"`
			FeeApr    float64 `json:"feeApr"`
			RewardApr struct {
				A float64 `json:"A"`
				B float64 `json:"B"`
				C float64 `json:"C"`
			} `json:"rewardApr"`
			Apr      float64 `json:"apr"`
			PriceMin float64 `json:"priceMin"`
			PriceMax float64 `json:"priceMax"`
		} `json:"week"`
		Month struct {
			Volume    float64 `json:"volume"`
			VolumeFee float64 `json:"volumeFee"`
			FeeA      float64 `json:"feeA"`
			FeeB      float64 `json:"feeB"`
			FeeApr    float64 `json:"feeApr"`
			RewardApr struct {
				A float64 `json:"A"`
				B float64 `json:"B"`
				C float64 `json:"C"`
			} `json:"rewardApr"`
			Apr      float64 `json:"apr"`
			PriceMin float64 `json:"priceMin"`
			PriceMax float64 `json:"priceMax"`
		} `json:"month"`
		LookupTableAccount string `json:"lookupTableAccount"`
	} `json:"data"`
}

type Token struct {
	Symbol     string `json:"symbol"`
	Name       string `json:"name"`
	Mint       string `json:"mint"`
	Decimals   int    `json:"decimals"`
	Extensions struct {
		CoingeckoID string `json:"coingeckoId"`
	} `json:"extensions,omitempty"`
	Icon      string `json:"icon"`
	HasFreeze int    `json:"hasFreeze"`
}

type TokenList struct {
	Name       string  `json:"name"`
	Official   []Token `json:"official"`
	UnOfficial []Token `json:"unOfficial"`
	UnNamed    []struct {
		Mint      string `json:"mint"`
		Decimals  int    `json:"decimals"`
		HasFreeze int    `json:"hasFreeze"`
	} `json:"unNamed"`
	Blacklist []interface{} `json:"blacklist"`
}

// find name and symbol
func (t TokenList) FindInfo(mintAsString string) Token {
	var token Token
	for _, item := range t.Official {
		if item.Mint == mintAsString {
			return token
		}
	}
	for _, item := range t.UnOfficial {
		if item.Mint == mintAsString {
			return token
		}
	}
	return token
}
