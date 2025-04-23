package dao

import (
	"context"
	"fmt"
	"navigator/model"
)

func (d *Dao) GetUserAll(ctx context.Context) ([]model.User, error) {
	result := make([]model.User, 0)
	err := d.mySqlSlave.Table("users").Find(&result).Error
	if err != nil {
		fmt.Println("GetUserAll error:", err)
		return nil, err
	}
	return result, nil
}
