package services

import (
	"blog/common"
	"blog/models"
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"log"
	"math"
)

func InsertAddress(addresses []models.FilAddresses) *gorm.DB {
	return common.DB.CreateInBatches(&addresses, 100)
}

func UpdateBalance(address models.FilAddresses) {

	tx := common.DB.Save(&address) //.Exec(sql)
	if tx != nil {
		log.Printf("update %s success:\n", address.Balance)
	}
}

func UpdateAddrRealCount(realCount int, accountId string) {
	sql := fmt.Sprintf("UPDATE fil_addresses SET real_count=%d where account_id='%s'", realCount, accountId)
	tx := common.DB.Exec(sql)
	if tx != nil {
		log.Printf("update account %s real_count success:\n", accountId)
	}
}

func FindAllAddress(account string) []models.FilAddresses {
	var db = common.DB
	var addresses []models.FilAddresses
	tx := db.Model(&models.FilAddresses{})

	whereStr := "status = 1"
	if len(account) > 0 {
		whereStr += " AND account_id = ?"
		tx.Where(whereStr, account)
	} else {
		tx.Where(whereStr)
	}

	tx.Order("id").Find(&addresses)

	return addresses
}

type AddressGroup []struct {
	Node   string  `json:"node"`
	Amount float64 `json:"amount"`
}

func GetAddressMap() map[string]decimal.Decimal {
	var result AddressGroup
	common.DB.Model(&models.FilAddresses{}).Select("node, SUM(burn_amount) as amount").Group("node").Find(&result)

	var m = make(map[string]decimal.Decimal)
	for _, rel := range result {
		m[rel.Node] = decimal.NewFromFloat(math.Abs(rel.Amount))
	}
	return m
}
