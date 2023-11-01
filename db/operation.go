package db

import (
	"fmt"

	"github.com/hjunior29/chain-processor/models"
	"gorm.io/gorm"
)

func SaveLog(db *gorm.DB, err error, message string) error {
	errString := ""
	if err != nil {
		errString = err.Error()
	}

	log := models.Logs{
		Error:   errString,
		Message: message,
	}

	if err := db.Create(&log).Error; err != nil {
		return fmt.Errorf("Failed to register log: %w", err)
	}

	return nil
}

func SaveValidProduct(db *gorm.DB, BlockNumber, TimeStamp, Hash, MethodId, FunctionName string) error {
	validProduct := models.ValidProducts{
		BlockNumber:  BlockNumber,
		TimeStamp:    TimeStamp,
		Hash:         Hash,
		MethodId:     MethodId,
		FunctionName: FunctionName,
	}

	if err := db.Create(&validProduct).Error; err != nil {
		return fmt.Errorf("Failed to register valid product: %w", err)
	}

	return nil
}
