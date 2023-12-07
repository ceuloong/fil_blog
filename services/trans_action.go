package services

import (
	"blog/common"
	"fmt"
)

func transActionUpdate(value interface{}) {

	var db = common.DB
	tx := db.Begin()
	if tx.Error != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		fmt.Printf("begin trans action failed, err:%v\n", tx.Error.Error())
		return
	}

	save := tx.Save(value)
	if save.Error != nil {
		_ = tx.Rollback()
		fmt.Printf("exec save.RowsAffected() failed, err:%v\n", save.Error)
	}
}
