package core

import (
	"strings"

	"github.com/hjunior29/chain-processor/db"
	"github.com/hjunior29/chain-processor/models"
	"gorm.io/gorm"
)

func Processor(tx []models.Transaction) error {
	for _, tx := range tx {
		if strings.Contains(tx.FunctionName, "createProduct") {
			var existingProduct models.ValidProducts

			if result := db.DB.First(&existingProduct, "block_number = ?", tx.BlockNumber).Error; result == gorm.ErrRecordNotFound {
				if err := db.SaveValidProduct(db.DB, tx.BlockNumber, tx.TimeStamp, tx.Hash, tx.MethodId, tx.FunctionName); err != nil {
					return err
				}
			}

			if err := db.SaveLog(db.DB, nil, "The transaction: " + tx.Hash + " has createProduct: " + tx.BlockNumber); err != nil {
				return err
			}

			// Send Notification
		}
	}
	return nil
}
