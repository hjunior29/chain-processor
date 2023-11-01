package core

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hjunior29/chain-processor/models"
)

func GetTransactions(API_URL string) (*models.ApiResponse, error) {
	resp, err := http.Get(API_URL)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request to %s: %v", API_URL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch transactions, status code: %d", resp.StatusCode)
	}

	var apiResponse models.ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return nil, err
	}

	return &apiResponse, nil
}
