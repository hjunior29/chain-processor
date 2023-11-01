package core

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hjunior29/chain-processor/db"
	"github.com/hjunior29/chain-processor/models"
)

func GetTransactions(API_URL string) (*models.ApiResponse, error) {
	resp, err := http.Get(API_URL)
	if err != nil {
		if err := db.SaveLog(db.DB, err, "Failed to send HTTP request"); err != nil {
			return nil, err
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if err := db.SaveLog(db.DB, err, "Failed to fetch transactions, status code: " + fmt.Sprint(resp.StatusCode)); err != nil {
			return nil, err
		}
	}

	var apiResponse models.ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		if err := db.SaveLog(db.DB, err, "Failed to decode reponse"); err != nil {
			return nil, err
		}
	}

	return &apiResponse, nil
}
