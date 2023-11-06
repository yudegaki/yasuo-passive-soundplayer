package lolapi

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// ref: https://developer.riotgames.com/docs/lol
const ALLGAMEDATA_API_ENDPOINT = "https://127.0.0.1:2999/liveclientdata/allgamedata"

type LoLAllGameData struct {
	ActivePlayer Summoner `json:"activePlayer"`
}

func GetAllGameData() (*Summoner, error){
	// skip tls verification
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, err := http.NewRequest("GET", ALLGAMEDATA_API_ENDPOINT, nil)
	if err != nil {
		return nil, errors.New("Failed to create http request")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("Failed to send http request")
	}
	defer resp.Body.Close()

	// Read the response body into a byte array
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Failed to read response body")
	}

	var lolAllGameData LoLAllGameData
	err = json.Unmarshal(body, &lolAllGameData)
	if err != nil {
		return nil, errors.New("Failed to unmarshal json")
	}
	return &lolAllGameData.ActivePlayer, nil
}
